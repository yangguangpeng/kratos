package global

import (
	"helloworld/internal/conf"
	"helloworld/pkg/global/envFlag"
	"os"
)

const AUTO_GENERATION_PATH = `_autoGeneration`
const LOG_PATH = `logs`

func Initial(baseConf conf.Base) {
	envFlag.Instance = envFlag.HandlerInstance(baseConf.GetEnv())
}

func GetAppPath() string {
	appPath := os.Getenv(`APP_PATH`)
	if appPath == `` {
		appPath, _ = os.Getwd()
	}
	return appPath
}
