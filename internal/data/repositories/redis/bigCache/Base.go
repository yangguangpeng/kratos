package bigCache

import (
	"helloworld/pkg/cache/toolRedis"
)

type Base struct {
	toolRedis.BaseRedis
}

func (t *Base) Init() {
	t.BaseRedis.RedisPool = toolRedis.GetPool("bigCache", 0)
}
