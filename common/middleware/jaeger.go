package middleware

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	opentracinglog "github.com/opentracing/opentracing-go/log"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/server"
	"strings"
)

func JaegerMiddleware(tracer opentracing.Tracer) server.HandlerWrapper {
	return func(handlerFunc server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			md, ok := metadata.FromContext(ctx)
			if !ok {
				md = metadata.Metadata{}
			}
			var opts []opentracing.StartSpanOption
			if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
				opts = append(opts, opentracing.ChildOf(parentSpan.Context()))
			} else if spanCtx, err := tracer.Extract(opentracing.TextMap, opentracing.TextMapCarrier(md)); err == nil {
				opts = append(opts, opentracing.ChildOf(spanCtx))
			}
			// allocate new map with only one element
			nmd := make(metadata.Metadata, 1)
			name := req.Service() + "-" + req.Method()
			sp := tracer.StartSpan(name, opts...)
			for k, v := range nmd {
				md.Set(strings.Title(k), v)
			}
			if err := tracer.Inject(sp.Context(), opentracing.TextMap, opentracing.TextMapCarrier(nmd)); err != nil {
				return handlerFunc(ctx, req, rsp)
			}
			defer sp.Finish()
			ctx = opentracing.ContextWithSpan(ctx, sp)
			ctx = metadata.NewContext(ctx, md)
			err := handlerFunc(ctx, req, rsp)
			if err != nil {
				ext.Error.Set(sp, true)
				sp.LogFields(opentracinglog.String(req.Service()+"的"+req.Method()+"err:", err.Error()))
			}
			return err
		}
	}
}
