package redis

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomodule/redigo/redis"
	"helloworld/internal/conf"
	"time"
)

func getPool(redisConfig *conf.Data_Redis)  (*redis.Pool ,error){
	pool := &redis.Pool{
		MaxIdle:     int(redisConfig.GetMaxIdle()),
		//MaxActive:   config.MaxActive,
		//IdleTimeout: time.Duration(config.IdleTimeout) * time.Second,
		//Wait:        config.Wait,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisConfig.GetAddr())
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	key := "___redis_pool_test_flag___"
	//log.Println("key:=", key)
	redis1 := pool.Get()
	_, err := redis1.Do("SET", key, time.Now().String())
	if err != nil {
		log.Log(log.LevelError , "redis error")
	}
	return pool, err
}
