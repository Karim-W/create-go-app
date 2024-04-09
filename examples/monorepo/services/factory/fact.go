package factory

import (
	"{{.moduleName}}/pkg/infrastructure/tracing"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/karim-w/stdlib"
	"github.com/karim-w/stdlib/httpclient"
	"github.com/karim-w/stdlib/sqldb"
	"github.com/soreing/trex"
	"go.uber.org/zap"
)

type Service interface {
	context.Context
	Logger() *zap.Logger
	Context() context.Context
	Trx() tracing.Tracer
	TraceParent() string
	Span() (string, string)
	Child() Service
	Caller() string
	HttpClient(url string) httpclient.HTTPRequest
	PSQL() sqldb.DB
	TraceInfo() (ver, tid, pid, rid, flg string)
	SetCaller(caller string)
	RDB() *redis.Client
	DLock(key string) (*redsync.Mutex, error)
}

type sf struct {
	traceparent string
	logger      *zap.Logger
	ctx         context.Context
	caller      string
}

func (s *sf) Child() Service {
	ver, tid, pid, rid, flg := deps.Trx.ExtractTraceInfo(s.ctx)

	ftx, _ := newFactoryFromTraceParentWithRid(
		ver+"-"+tid+"-"+pid+"-"+flg,
		rid,
	)

	return ftx
}

func (s *sf) DLock(key string) (*redsync.Mutex, error) {
	s.logger.Info("acquiring lock", zap.String("key", key))
	mtx := deps.RSync.NewMutex(key)

	err := mtx.Lock()
	if err != nil {
		s.logger.Error("failed to acquire lock", zap.Error(err))
		return nil, err
	}

	s.logger.Info("lock acquired", zap.String("key", key))
	return mtx, nil
}

func (s *sf) RDB() *redis.Client {
	return deps.Redis
}

func (s *sf) SetCaller(caller string) {
	s.caller = caller
}

func (s *sf) TraceInfo() (ver string, tid string, pid string, rid string, flg string) {
	return deps.Trx.ExtractTraceInfo(s.ctx)
}

func (s *sf) PSQL() sqldb.DB {
	return deps.PSQL
}

func (s *sf) Trx() tracing.Tracer {
	return deps.Trx
}

// Span()
// returns a new traceparent based on the current traceparent
func (s *sf) Span() (string, string) {
	ver, tid, _, rid, flg := deps.Trx.ExtractTraceInfo(s.ctx)
	sid, _ := stdlib.GenerateParentId()
	if sid == "" {
		sid = rid
	}

	return ver + "-" + tid + "-" + sid + "-" + flg, sid
}

func (s *sf) Caller() string {
	return s.caller
}

func NewFactory(ctx context.Context) Service {
	ver, tid, pid, rid, flg := deps.Trx.ExtractTraceInfo(ctx)
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
	ctx := context.WithValue(context.Background(), "tinfo", txm)
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
	ctx := context.WithValue(context.Background(), "tinfo", txm)
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
	ver, tid, _, rid, flg := deps.Trx.ExtractTraceInfo(s.ctx)
	if sid == "" {
		sid = rid
	}
	reqTraceparent := fmt.Sprintf("%s-%s-%s-%s", ver, tid, sid, flg)
	r := httpclient.Req(url).
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
				s.Context(),
				sid,
				"http",
				hostName,
				url,
				err == nil && statusCode < 400 && statusCode >= 200,
				meta.StartTime,
				meta.EndTime,
				m,
			)
			s.Logger().Info("Httpclient AfterHook",
				zap.String("traceparent", reqTraceparent),
				zap.String("url", url),
				zap.String("host", hostName),
				zap.String("method", req.Method),
				zap.String("status", status),
			)
			return
		}).AddHeader("traceparent", reqTraceparent).
		AddHeader("Content-Type", "application/json")
	return r
}
