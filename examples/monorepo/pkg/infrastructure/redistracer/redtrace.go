package redistracer

import (
	"{{.moduleName}}/services/factory"
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/karim-w/stdlib"
)

type TraceHook struct {
	serviceName string
	tracer      Tracer
}

func WrapWithTracing(
	rd *redis.Client,
	serviceName string,
	tracer Tracer,
) *redis.Client {
	rd.AddHook(&TraceHook{
		serviceName: serviceName,
		tracer:      tracer,
	})

	return rd
}

var _ redis.Hook = (*TraceHook)(nil)

const START_KEY = "t_start"

func (h *TraceHook) BeforeProcess(
	ctx context.Context,
	cmd redis.Cmder,
) (context.Context, error) {
	return context.WithValue(ctx, START_KEY, time.Now()), nil
}

func (h *TraceHook) AfterProcess(
	ctx context.Context,
	cmd redis.Cmder,
) error {
	end := time.Now()
	start := end

	ftx := factory.NewFactory(ctx)

	raw := ctx.Value(START_KEY)

	if raw != nil {
		if cstd, ok := raw.(time.Time); ok {
			start = cstd
		}
	}

	fields := map[string]string{}

	err := cmd.Err()
	if err != nil {
		fields["error"] = err.Error()
	}

	sid, _ := stdlib.GenerateParentId()
	if sid == "" {
		_, _, _, rid, _ := h.tracer.ExtractTraceInfo(ctx)
		sid = rid
	}

	ftx.Trx().TraceDependency(
		ctx,
		sid,
		"Redis",
		h.serviceName,
		cmd.String(),
		err == nil,
		start,
		end,
		fields,
	)

	return nil
}

func (h *TraceHook) BeforeProcessPipeline(
	ctx context.Context,
	cmds []redis.Cmder,
) (context.Context, error) {
	return context.WithValue(ctx, START_KEY, time.Now()), nil
}

func (h *TraceHook) AfterProcessPipeline(
	ctx context.Context,
	cmds []redis.Cmder,
) error {
	end := time.Now()
	start := end

	raw := ctx.Value(START_KEY)
	if raw != nil {
		if cstd, ok := raw.(time.Time); ok {
			start = cstd
		}
	}

	fields := map[string]string{}
	cmdstr := ""

	var err error

	for i := 0; i < len(cmds); i++ {
		er := cmds[i].Err()
		if er != nil {
			err = errors.Join(err, er)
		}
		cmdstr = cmdstr + cmds[i].String() + ";"
	}

	if err != nil {
		fields["error"] = err.Error()
	}

	sid, _ := stdlib.GenerateParentId()
	if sid == "" {
		_, _, _, rid, _ := h.tracer.ExtractTraceInfo(ctx)
		sid = rid
	}

	h.tracer.TraceDependency(
		ctx,
		sid,
		"Redis",
		h.serviceName,
		cmdstr,
		len(fields) == 0,
		start,
		end,
		fields,
	)
	return nil
}
