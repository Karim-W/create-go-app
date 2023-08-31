package factory

import (
	"context"

	"{{.moduleName}}/internal/constants"
	"{{.moduleName}}/pkg/infra/tracing"

	"github.com/soreing/trex"
	"go.uber.org/zap"
)

type depenencies struct {
	logger *zap.Logger
	trx    tracing.Tracer
}

var deps *depenencies

// not thread safe
func SetUpDependencies(logger *zap.Logger, trx tracing.Tracer) {
	if deps != nil {
		return
	}
	deps = &depenencies{
		logger, trx,
	}
}

type Service interface {
	Logger() *zap.Logger
	Context() context.Context
	TraceParent() string
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
		logger:      deps.logger.With(zap.String("traceparent", traceparent)),
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
		logger:      deps.logger.With(zap.String("traceparent", traceparent)),
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
