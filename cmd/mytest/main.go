package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/syyongx/php2go"
	"os"
)

func main() {

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

	fmt.Println(logFile)

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
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	h := log.NewHelper(logger)

	h.Info(`sdfsdfsd`)
}

func GetT() log.Valuer {
	return func(ctx context.Context) interface{} {
		if header, ok := transport.FromServerContext(ctx); ok {
			return header.RequestHeader().Get(`myTraceID`)
		}
		return "nothing"
	}
}
