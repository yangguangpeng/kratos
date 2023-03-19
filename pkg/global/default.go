package global

import (
	"helloworld/internal/conf"
	"helloworld/pkg/global/envFlag"
	toolSentry "helloworld/pkg/sentry"
)

const AUTO_GENERATION_PATH = `_autoGeneration`
const LOG_PATH = `logs`

func Initial(baseConf conf.Base) {
	envFlag.Instance = envFlag.HandlerInstance(baseConf.GetEnv())
	// sentry, 所有go框架共用同一个dsn
	toolSentry := &toolSentry.InitSentry{
		Dsn: `conf.Sentry.Dsn`,
	}
	toolSentry.Init(`conf.ProjectName`)
	defer toolSentry.Quit()
}
