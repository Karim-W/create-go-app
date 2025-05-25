package factory

import (
	"{{.moduleName}}/pkg/infrastructure/tracing"

	"github.com/karim-w/stdlib"
)

func (s *sf) SetCaller(caller string) {
	s.caller = caller
}

func (s *sf) Caller() string {
	return s.caller
}

func (s *sf) TraceInfo() (ver string, tid string, pid string, rid string, flg string) {
	return deps.Trx.ExtractTraceInfo(s.ctx)
}

func (s *sf) Trx() tracing.Tracer {
	return deps.Trx
}

// TraceParent()
// returns the traceparent
func (s *sf) TraceParent() string {
	return s.traceparent
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
