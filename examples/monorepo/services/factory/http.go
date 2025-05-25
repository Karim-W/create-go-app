package factory

import (
	"net/http"
	"strings"

	"github.com/karim-w/stdlib"
	"github.com/karim-w/stdlib/httpclient"
	"go.uber.org/zap"
)

func (s *sf) HttpClient(url string) httpclient.HTTPRequest {
	sid, _ := stdlib.GenerateParentId()

	if sid == "" {
		sid = s.trace.Rid
	}

	reqTraceparent := s.trace.Ver + "-" + s.trace.Tid + "-" + sid + "-" + s.trace.Flg

	r := httpclient.Req(url).
		AddAfterHook(func(
			req *http.Request,
			res *http.Response,
			meta httpclient.HTTPMetadata,
			err error,
		) {
			span_id := req.Header.Get("x-span-id")
			req_traceparent := req.Header.Get("traceparent")

			ftx, e := NewFactoryFromTraceParent(req_traceparent)
			if e != nil {
				return
			}

			if span_id == "" {
				span_id, _ = stdlib.GenerateParentId()
			}

			hostName := req.URL.Hostname()
			urlbuilder := strings.Builder{}
			urlbuilder.WriteString(req.Method)
			urlbuilder.WriteString(" ")
			urlbuilder.WriteString(req.URL.RequestURI())
			url := urlbuilder.String()
			m := map[string]string{
				"host": hostName,
				"url":  url,
			}
			if err != nil {
				m["error"] = err.Error()
			}
			status := "unknown"
			statusCode := 500
			if res != nil {
				status = res.Status
				statusCode = res.StatusCode
			}

			deps.Trx.TraceDependency(
				ftx.Context(),
				span_id,
				"http",
				hostName,
				url,
				err == nil && statusCode < 400 && statusCode >= 200,
				meta.StartTime,
				meta.EndTime,
				m,
			)

			ftx.Logger().Info("Httpclient AfterHook",
				zap.String("traceparent", reqTraceparent),
				zap.String("url", url),
				zap.String("host", hostName),
				zap.String("method", req.Method),
				zap.String("status", status),
			)

			return
		}).AddHeader("traceparent", reqTraceparent).
		AddHeader("x-span-id", sid).
		AddHeader("x-caller-id", s.Caller())

	return r
}
