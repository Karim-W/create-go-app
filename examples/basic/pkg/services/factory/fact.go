package factory

import (
	"context"
	"fmt"
	"strings"

	"{{.moduleName}}/internal/constants"
	"{{.moduleName}}/pkg/infra/logger"
	"{{.moduleName}}/pkg/infra/tracing"

	"github.com/karim-w/stdlib"
	"github.com/soreing/trex"
	"go.uber.org/zap"
)

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
	// Clean up the context from the request context and making transient deps
	if ctx == nil {
		ctx = context.TODO()
	}
	traceparent := ""
	// check if tid is already in the context
	raw := ctx.Value(constants.TRACE_INFO_KEY)
	if raw == nil {
		tid, err := stdlib.GenerateNewTraceparent(true)
		ver, tid, pid, flg, _ := trex.DecodeTraceparent(tid)
		// Generate a new resource id
		rid, err := trex.GenerateRadomHexString(8)

		if err == nil {
			// Generate a transaction context usin the factory
			txm := trex.TxModel{
				Ver: ver,
				Tid: tid,
				Pid: pid,
				Rid: rid,
				Flg: flg,
			}
			traceparentBuilder := strings.Builder{}
			traceparentBuilder.WriteString(ver)
			traceparentBuilder.WriteString("-")
			traceparentBuilder.WriteString(tid)
			traceparentBuilder.WriteString("-")
			traceparentBuilder.WriteString(pid)
			traceparentBuilder.WriteString("-")
			traceparentBuilder.WriteString(rid)
			traceparentBuilder.WriteString("-")
			traceparentBuilder.WriteString(flg)
			traceparent = traceparentBuilder.String()
			ctx = context.WithValue(ctx, constants.TRACE_INFO_KEY, txm)
		}
	}

	return &sf{
		logger:      logger.GetTracedLogger(traceparent),
		ctx:         ctx,
		traceparent: traceparent,
	}
}

func NewFactoryFromTraceParent(traceparent string) (Service, error) {
	if traceparent == "" {
		return NewFactory(nil), nil
	}
	ver, tid, pid, flg, err := trex.DecodeTraceparent(traceparent)
	if err != nil {
		return nil, err
	}
	// Generate a new resource id
	rid, err := trex.GenerateRadomHexString(8)
	if err != nil {
		return nil, err
	}
	// Generate a transaction context usin the factory
	txm := trex.TxModel{
		Ver: ver,
		Tid: tid,
		Pid: pid,
		Rid: rid,
		Flg: flg,
	}
	ctx := context.WithValue(context.Background(), constants.TRACE_INFO_KEY, txm)
	return &sf{
		logger:      logger.GetTracedLogger(traceparent),
		ctx:         ctx,
		traceparent: traceparent,
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
	var id string
	sid, err := stdlib.GenerateParentId()
	ver, tid, _, rid, flg := tracing.GetTracer().ExtractTraceInfo(s.ctx)
	if err != nil {
		id = rid
	} else {
		id = sid
	}
	return fmt.Sprintf("%s-%s-%s-%s", ver, tid, id, flg)
}
