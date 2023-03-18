package toolRedis

import (
	"log"
	"time"
)

type InitRedis struct {
	RedisInfo           RedisSchema
	SystemQuit          chan int
	HealthCheckDuration int
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
			SetRedisPool(connectionInfo, redisName, index)
		}
	}

	go config.HealthCheck()
}

func (config *InitRedis) Quit() {
	config.SystemQuit <- 1
	log.Println("InitRedis.Quit() 退出成功")
}

func (config *InitRedis) Close() {

	for redisName, connectionInfos := range config.RedisInfo {
		for index, _ := range connectionInfos {
			CloseRedisPool(redisName, index)
		}
	}
	log.Println("InitRedis.Close exited")
}

// HealthCheck 健康检查
func (config *InitRedis) HealthCheck() {

	for redisName, connectionInfos := range config.RedisInfo {
		for index, _ := range connectionInfos {
			Reconnection(redisName, index)
		}
	}

	//定时检查
	select {
	case <-config.SystemQuit:
		config.Close()
		log.Println("InitRedis.healthCheck exited")
		return

	case <-time.After(time.Duration(config.HealthCheckDuration) * time.Millisecond):
		config.HealthCheck()
	}
}
