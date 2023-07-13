package factory

import (
	"context"
	"fmt"
	"{{.moduleName}}/internals/constants"
	"{{.moduleName}}/pkg/infra/logger"
	"{{.moduleName}}/pkg/infra/tracing"
	"strings"

	"github.com/karim-w/stdlib"
	"github.com/soreing/trex"
	"go.uber.org/zap"
)

type Service interface {
	GetLogger() *zap.Logger
	GetContext() context.Context
	GetTraceParent() string
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
		logger.GetLogger().
			Info("Creating New Context", zap.String("traceparent", tid))
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

// GetLogger() Returns the logger with the traceinfo
func (s *sf) GetLogger() *zap.Logger {
	s.logger.Info("Getting Logger")
	return s.logger
}

// GetContext() Returns the context of the request
func (s *sf) GetContext() context.Context {
	s.logger.Info("Getting Context")
	return s.ctx
}

// GetTraceParent()
// returns the traceparent
func (s *sf) GetTraceParent() string {
	s.logger.Info("Getting TraceParent")
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
