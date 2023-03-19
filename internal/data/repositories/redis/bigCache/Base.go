package bigCache

import (
	"helloworld/pkg/cache/toolRedis"
)

type Base struct {
	toolRedis.BaseRedis
}

func (t *Base) Init() {
	t.BaseRedis.RedisPool = toolRedis.GetPool("bigCacheMaster", 0)

}

func (t *Base) UseSlave() {
	t.BaseRedis.RedisPool = toolRedis.GetPool("bigCacheSlave", 0)
}
