package toolRedis

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomodule/redigo/redis"
	"sync"
)

var (
	allRedisPools   = map[string][]*RedisPool{}
	mutexRedisPools sync.RWMutex
)

// GetRedisPool 对外可用的地址
func GetRedisPool(redisName string, key int) *RedisPool {
	mutexRedisPools.RLock()
	defer mutexRedisPools.RUnlock()

	if _, ok := allRedisPools[redisName]; !ok {
		return nil
	}

	if len(allRedisPools[redisName]) <= key {
		return nil
	}

	return allRedisPools[redisName][key]
}

func GetPool(redisName string, key int) *redis.Pool {
	mutexRedisPools.RLock()
	defer mutexRedisPools.RUnlock()

	if _, ok := allRedisPools[redisName]; !ok {
		return nil
	}

	if len(allRedisPools[redisName]) <= key {
		return nil
	}

	return allRedisPools[redisName][key].Pool()
}

func InitRedisPool(redisName string, length int) {
	mutexRedisPools.Lock()
	defer mutexRedisPools.Unlock()

	allRedisPools[redisName] = make([]*RedisPool, length)
}

func SetRedisPool(config RedisItemSchema, redisName string, key int, log *log.Helper) {
	mutexRedisPools.Lock()
	defer mutexRedisPools.Unlock()

	allRedisPools[redisName][key] = &RedisPool{Log: log}

	allRedisPools[redisName][key].Init(config, redisName, key)
}

func CloseRedisPool(redisName string, key int) {
	mutexRedisPools.Lock()
	defer mutexRedisPools.Unlock()

	allRedisPools[redisName][key].Close()
}

func Reconnection(redisName string, key int) {
	mutexRedisPools.Lock()
	defer mutexRedisPools.Unlock()

	allRedisPools[redisName][key].Connect()
}
