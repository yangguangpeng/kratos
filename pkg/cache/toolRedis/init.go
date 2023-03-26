package toolRedis

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type InitRedis struct {
	RedisInfo           RedisSchema
	SystemQuit          chan int
	HealthCheckDuration int
	Log                 *log.Helper
}

func (config *InitRedis) Init() {

	if config.HealthCheckDuration < 1000 {
		config.HealthCheckDuration = 5000
	}

	config.SystemQuit = make(chan int, 1)
	//RedisPools = map[string][]*redis.Pool{}

	for redisName, connectionInfos := range config.RedisInfo {
		//RedisPools[redisName] = make([]*redis.Pool, len(connectionInfos))
		InitRedisPool(redisName, len(connectionInfos))

		for index, connectionInfo := range connectionInfos {
			SetRedisPool(connectionInfo, redisName, index, config.Log)
		}
	}

	go config.HealthCheck()
}

func (config *InitRedis) Quit() {
	fmt.Println(`redis quit...`)
	config.SystemQuit <- 1
	config.Log.Info("InitRedis.Quit() 退出成功")
}

func (config *InitRedis) Close() {

	for redisName, connectionInfos := range config.RedisInfo {
		for index, _ := range connectionInfos {
			CloseRedisPool(redisName, index)
		}
	}
	config.Log.Info("InitRedis.Close exited")
}

// HealthCheck 健康检查
func (config *InitRedis) HealthCheck() {

	fmt.Println(`HealthCheck.....`)
	for redisName, connectionInfos := range config.RedisInfo {
		for index, _ := range connectionInfos {
			Reconnection(redisName, index)
		}
	}

	//定时检查
	select {
	case <-config.SystemQuit:
		config.Close()
		config.Log.Info("InitRedis.healthCheck exited")
		return

	case <-time.After(time.Duration(config.HealthCheckDuration) * time.Millisecond):
		config.HealthCheck()
	}
}
