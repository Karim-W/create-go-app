package factory

import (
	"context"
	"sync"

	"github.com/karim-w/stdlib/sqldb"
	"go.uber.org/zap"
)

type tracectx struct {
	Ver, Tid, Pid, Rid, Flg string
}

type sf struct {
	traceparent string
	logger      *zap.Logger
	ctx         context.Context
	caller      string
	psqlTx      *sqldb.Tx
	sql_mtx     *sync.RWMutex
	trace       tracectx
	store       *sync.Map
}
