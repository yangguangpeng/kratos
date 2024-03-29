package globalLog

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/syyongx/php2go"
	"helloworld/pkg/global"
	"helloworld/pkg/global/envFlag"
	"os"
)

func New() (logger log.Logger, err error) {

	stdout := os.Stdout
	if envFlag.Instance.IsEnvPro() {
		stdout, err = getLogFile()
	}
	//stdout, err = getLogFile()
	if err != nil {
		return nil, err
	}
	logger = log.With(log.NewStdLogger(stdout),
		"ts", log.DefaultTimestamp,
		"defaultCaller", log.DefaultCaller,
		"caller", log.Caller,
		"service.id", 1,
		"service.name", 1,
		"service.version", 1,
		"trace.id", TraceID(),
		//"span.id", tracing.SpanID(),
	)
	return
}

func getLogFile() (*os.File, error) {

	dir := global.GetAppPath()
	time := php2go.Time()
	logPath := dir + `/` + global.AUTO_GENERATION_PATH + `/` +
		global.LOG_PATH + `/` +
		php2go.Date(`2006-01`, time) + `/`
	if ok, _ := php2go.IsDir(logPath); !ok {
		err := os.MkdirAll(logPath, 0666)
		if err != nil {
			return nil, err
		}
	}
	logFileNameSuffix := `_` + php2go.Date(`02`, time)
	logFileName := `test` + logFileNameSuffix + `.log`
	logFile := logPath + `/` + logFileName
	//文件日志：
	return os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}

func TraceID() log.Valuer {
	return func(ctx context.Context) interface{} {
		if header, ok := transport.FromServerContext(ctx); ok {
			return header.RequestHeader().Get(`trace_id`)
		}
		return ``
	}
}
