package applicationinsightstrace

import (
	"context"
	"{{.moduleName}}/internals/constants"

	trace "github.com/BetaLixT/appInsightsTrace"
	"github.com/soreing/trex"
	"go.uber.org/zap"
)

type DefaultTraceExtractor struct{}

func (*DefaultTraceExtractor) ExtractTraceInfo(
	ctx context.Context,
) (ver, tid, pid, rid, flg string) {
	if tinfo, ok := ctx.Value(constants.TRACE_INFO_KEY).(trex.TxModel); !ok {
		return "", "", "", "", ""
	} else {
		return tinfo.Ver, tinfo.Tid, tinfo.Pid, tinfo.Rid, tinfo.Flg
	}
}

func InitOrDie(
	instrumentationKey string,
	serviceName string,
) *trace.AppInsightsCore {
	return trace.NewAppInsightsCore(&trace.AppInsightsOptions{
		ServiceName:        serviceName,
		InstrumentationKey: instrumentationKey,
	}, &DefaultTraceExtractor{}, zap.NewNop())
}
