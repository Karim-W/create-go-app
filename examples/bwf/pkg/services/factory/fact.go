package factory

import (
	"context"
	"net/http"
	"{{.moduleName}}/internal/constants"
	"strings"

	"github.com/karim-w/stdlib"
	"github.com/karim-w/stdlib/httpclient"
	"github.com/soreing/trex"
	"go.uber.org/zap"
)

type Service interface {
	Logger() *zap.Logger
	Context() context.Context
	TraceParent() string
	HttpClient(url string) httpclient.HTTPRequest
}

type sf struct {
	traceparent string
	logger      *zap.Logger
	ctx         context.Context
}

func NewFactory(ctx context.Context) Service {
	ver, tid, pid, rid, flg := deps.trx.ExtractTraceInfo(ctx)
	ftx, _ := newFactoryFromTraceParentWithRid(
		ver+"-"+tid+"-"+pid+"-"+flg,
		rid,
	)
	return ftx
}

func newFactoryFromTraceParentWithRid(traceparent string, rid string) (Service, error) {
	ver, tid, pid, flg, err := trex.DecodeTraceparent(traceparent)
	// If the header could not be decoded, generate a new header
	if err != nil {
		ver, flg = "00", "01"
		tid, _ = trex.GenerateRadomHexString(16)
		pid, _ = trex.GenerateRadomHexString(8)
	}

	// Generate a new resource id
	if rid == "" {
		rid, _ = trex.GenerateRadomHexString(8)
	}

	// Generate a transaction context usin the factory
	txm := trex.TxModel{
		Ver: ver,
		Tid: tid,
		Pid: pid,
		Rid: rid,
		Flg: flg,
	}

	TraceParent := ver + "-" + tid + "-" + pid + "-" + flg

	ctx := context.WithValue(context.Background(), constants.TRACE_INFO_KEY, txm)

	return &sf{
		logger:      zap.L().With(zap.String("traceparent", TraceParent)),
		ctx:         ctx,
		traceparent: TraceParent,
	}, nil
}

func NewFactoryFromTraceParent(traceparent string) (Service, error) {
	ver, tid, pid, flg, err := trex.DecodeTraceparent(traceparent)
	// If the header could not be decoded, generate a new header
	if err != nil {
		ver, flg = "00", "01"
		tid, _ = trex.GenerateRadomHexString(16)
		pid, _ = trex.GenerateRadomHexString(8)
	}
	// Generate a new resource id
	rid, _ := trex.GenerateRadomHexString(8)
	// Generate a transaction context usin the factory
	txm := trex.TxModel{
		Ver: ver,
		Tid: tid,
		Pid: pid,
		Rid: rid,
		Flg: flg,
	}
	TraceParent := ver + "-" + tid + "-" + pid + "-" + flg
	ctx := context.WithValue(context.Background(), constants.TRACE_INFO_KEY, txm)
	return &sf{
		logger:      zap.L().With(zap.String("traceparent", TraceParent)),
		ctx:         ctx,
		traceparent: TraceParent,
	}, nil
}

// Logger() Returns the logger with the traceinfo
func (s *sf) Logger() *zap.Logger {
	return s.logger
}

// Context() Returns the context of the request
func (s *sf) Context() context.Context {
	return s.ctx
}

// TraceParent()
// returns the traceparent
func (s *sf) TraceParent() string {
	return s.traceparent
}

// HTTPClient()
// returns the http client
func (s *sf) HttpClient(url string) httpclient.HTTPRequest {
	sid, _ := stdlib.GenerateParentId()

	ver, tid, _, rid, flg := deps.trx.ExtractTraceInfo(s.ctx)

	if sid == "" {
		sid = rid
	}

	reqBuffer := make([]byte, 0, len(ver)+len(tid)+len(sid)+len(flg)+4)

	reqBuffer = append(reqBuffer, ver...)
	reqBuffer = append(reqBuffer, '-')
	reqBuffer = append(reqBuffer, tid...)
	reqBuffer = append(reqBuffer, '-')
	reqBuffer = append(reqBuffer, sid...)
	reqBuffer = append(reqBuffer, '-')
	reqBuffer = append(reqBuffer, flg...)

	reqTraceparent := string(reqBuffer)

	return httpclient.Req(url).
		AddAfterHook(func(
			req *http.Request,
			res *http.Response,
			meta httpclient.HTTPMetadata,
			err error,
		) {
			hostName := req.URL.Hostname()

			urlbuilder := strings.Builder{}
			urlbuilder.WriteString(req.Method)
			urlbuilder.WriteString(" ")
			urlbuilder.WriteString(req.URL.RequestURI())
			url := urlbuilder.String()

			traceparent := req.Header.Get("traceparent")
			spanId := req.Header.Get("span-id")

			m := map[string]string{
				"host": hostName,
				"url":  url,
			}

			if err != nil {
				m["error"] = err.Error()
			}

			// Default status code to 500 and status to unknown
			status := "unknown"
			statusCode := 500

			if res != nil {
				status = res.Status
				statusCode = res.StatusCode
			}

			deps.trx.TraceDependency(
				s.Context(),
				spanId,
				"http",
				hostName,
				url,
				err == nil && statusCode < 400 && statusCode >= 200,
				meta.StartTime,
				meta.EndTime,
				m,
			)

			s.Logger().Info("Httpclient AfterHook",
				zap.String("traceparent", traceparent),
				zap.String("url", url),
				zap.String("host", hostName),
				zap.String("method", req.Method),
				zap.String("status", status),
			)

			return
		}).JSON().AddHeader("traceparent", reqTraceparent).
		AddHeader("span-id", sid)
}
