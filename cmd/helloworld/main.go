package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/syyongx/php2go"
	"helloworld/internal/conf"
	"helloworld/pkg/apollo"
	"os"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var baseConfig conf.Base
	if err := c.Scan(&baseConfig); err != nil {
		panic(err)
	}

	globalConfig := apollo.GetConfig(baseConfig.GetApollo())

	//日志开始
	time := php2go.Time()
	logPath := `_autoGeneration/logs/` + php2go.Date(`2006-01`, time)
	if ok, _ := php2go.IsDir(logPath); !ok {
		err := os.MkdirAll(logPath, 0666)
		if err != nil {
			panic(err)
		}
	}
	logFileNameSuffix := `_` + php2go.Date(`02`, time)
	logFileName := `test` + logFileNameSuffix + `.log`
	logFile := logPath + `/` + logFileName
	//文件日志：
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	logger := log.With(log.NewStdLogger(f),
		"ts", log.DefaultTimestamp,
		"defaultCaller", log.DefaultCaller,
		"caller", log.Caller,
		"service.id", 1,
		"service.name", 1,
		"service.version", 1,
		"trace.id", GetTraceID(),
		"span.id", tracing.SpanID(),
	)
	//日志结束

	app, cleanup, err := wireApp(baseConfig.Server, globalConfig, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	//start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func GetTraceID() log.Valuer {
	return func(ctx context.Context) interface{} {
		if header, ok := transport.FromServerContext(ctx); ok {
			return header.RequestHeader().Get(`MyTraceID`)
		}
		if md, ok := metadata.FromServerContext(ctx); ok {
			return md.Get("MyTraceID")
		}
		return "nothing"
	}
}
