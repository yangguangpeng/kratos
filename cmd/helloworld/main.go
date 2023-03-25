package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"helloworld/internal/conf"
	"helloworld/pkg/apollo"
	"helloworld/pkg/global"
	"helloworld/pkg/global/globalLog"
	toolSentry "helloworld/pkg/sentry"
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

	logger, logerror := globalLog.New()
	if logerror != nil {
		panic(logerror)
	}

	global.Initial(baseConfig)

	// sentry, 所有go框架共用同一个dsn
	toolSentry := &toolSentry.InitSentry{
		Dsn: `conf.Sentry.Dsn`,
	}
	toolSentry.Init(`conf.ProjectName`)
	defer toolSentry.Quit()

	app, cleanup, err := wireApp(baseConfig.Server, globalConfig, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	//start and wait for stop signal
	if err = app.Run(); err != nil {
		panic(err)
	}
}
