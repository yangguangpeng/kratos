package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/conf"
	"helloworld/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, demo *service.DemoService, logger log.Logger) *grpc.Server {
	err := initTracer("http://localhost:14268/api/traces")
	if err != nil {
		panic(err)
	}
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			//jwt验证中间件
			//jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
			//	return []byte("testKey"), nil
			//}),
			MyMiddleware(),
			tracing.Server(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	v1.RegisterDemoServer(srv, demo)
	return srv
}

// Set global trace provider
func initTracer(url string) error {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return err
	}
	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the parent span to 100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String("kratos-trace"),
			attribute.String("exporter", "jaeger"),
			attribute.Float64("float", 312.23),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}

//MyMiddleware 自定义中间件
func MyMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			log.Info(`MyMiddleware 开始`)

			if ServerContext, ok := transport.FromServerContext(ctx); ok {
				authorization := ServerContext.RequestHeader().Get(`Authorization`)
				log.Infow(`MyMiddleware print: 请求的认证信息`, authorization)
			}
			log.Info(`MyMiddleware 结束`)

			return handler(ctx, req)
		}
	}
}
