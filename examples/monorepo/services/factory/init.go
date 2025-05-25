package factory

import (
	"context"
	"sync"

	"github.com/karim-w/stdlib"
	"github.com/soreing/trex"
	"go.uber.org/zap"
)

func NewFactory(ctx context.Context) Service {
	ver, tid, pid, _, flg := deps.Trx.ExtractTraceInfo(ctx)
	if tid == "" || pid == "" || ver == "" || flg == "" {
		ftx, _ := NewFactoryFromTraceParent("")
		return ftx
	}

	rid, _ := trex.GenerateRadomHexString(8)

	txm := tracectx{
		Ver: ver,
		Tid: tid,
		Pid: pid,
		Rid: rid,
		Flg: flg,
	}

	traceparent := ver + "-" + tid + "-" + rid + "-" + flg

	return &sf{
		traceparent: traceparent,
		logger:      zap.L().With(zap.String("traceparent", traceparent)),
		ctx:         ctx,
		trace:       txm,
		sql_mtx:     &sync.RWMutex{},
		store:       &sync.Map{},
	}
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

	sid, _ := stdlib.GenerateParentId()
	if rid == "" {
		rid = sid
	}

	// Generate a transaction context usin the factory
	txm := tracectx{
		Ver: ver,
		Tid: tid,
		Pid: pid,
		Rid: rid,
		Flg: flg,
	}

	TraceParent := ver + "-" + tid + "-" + rid + "-" + flg

	return &sf{
		logger:      zap.L().With(zap.String("traceparent", TraceParent)),
		ctx:         context.TODO(),
		traceparent: TraceParent,
		trace:       txm,
		sql_mtx:     &sync.RWMutex{},
		store:       &sync.Map{},
	}, nil
}

func (s *sf) Child() Service {
	ftx, err := NewFactoryFromTraceParent(s.traceparent)
	if err != nil {
		return s
	}
	return ftx
}
