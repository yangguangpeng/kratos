package global

import (
	"helloworld/internal/conf"
	"helloworld/pkg/global/envFlag"
)

func Initial(baseConf conf.Base) {
	envFlag.Instance = envFlag.HandlerInstance(baseConf.GetEnv())
}
