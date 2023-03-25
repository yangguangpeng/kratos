package bigCache

import (
	"helloworld/pkg/cache"
	"helloworld/pkg/cache/toolRedis"
)

type Base struct {
	toolRedis.BaseRedis
}

func (t *Base) Init() {
	t.BaseRedis.RedisPool = toolRedis.GetPool(cache.BIG_CACHE_MASTER, 0)
}

func (t *Base) UseSlave() {
	t.BaseRedis.RedisPool = toolRedis.GetPool(cache.BIG_CACHE_MASTER, 0)
}
