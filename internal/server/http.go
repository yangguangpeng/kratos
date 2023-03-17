package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	sentrykratos "github.com/go-kratos/sentry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/conf"
	"helloworld/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, Demo *service.DemoService, logger log.Logger) *http.Server {
	err := initTracer("http://192.168.1.5:14268/api/traces")
	if err != nil {
		panic(err)
	}
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			sentrykratos.Server(), // must after Recovery middleware, because of the exiting order will be reversed
			tracing.Server(),
			MyMiddleware1(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	v1.RegisterDemoHTTPServer(srv, Demo)
	return srv
}

func MyMiddleware1() middleware.Middleware {

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

			tp := otel.GetTracerProvider()

			tid, err := trace.TraceIDFromHex("abc")
			sid, err := trace.SpanIDFromHex("EF")
			sc := trace.NewSpanContext(trace.SpanContextConfig{TraceID: tid, SpanID: sid})

			ctx = trace.ContextWithSpanContext(ctx, sc)
			ctx, span := tp.Tracer("test instrumentation").Start(ctx, "span1")
			span.SetName(`hello`)
			var attrs []attribute.KeyValue
			attrs = append(attrs, semconv.HTTPMethodKey.String(`aaaa`))
			span.SetAttributes(attrs...)

			span.End()
			log.Info(`MyMiddleware1 开始`)

			log.Info(`MyMiddleware1 结束`)

			return handler(ctx, req)
		}
	}
}
