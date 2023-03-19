package global

import (
	"helloworld/internal/conf"
	"helloworld/pkg/global/envFlag"
)

const AUTO_GENERATION_PATH = `_autoGeneration`
const LOG_PATH = `logs`

func Initial(baseConf conf.Base) {
	envFlag.Instance = envFlag.HandlerInstance(baseConf.GetEnv())
}
