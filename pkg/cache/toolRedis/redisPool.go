package toolRedis

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomodule/redigo/redis"
	toolSentry "helloworld/pkg/sentry"
	"strconv"
	"time"
)

var (
	maxExpireSeconds = 10
	maxRetryTimes    = 3
)

type RedisPool struct {
	pool      *redis.Pool
	config    RedisItemSchema
	redisName string
	seq       int
	Log       *log.Helper
}

func (t *RedisPool) Pool() *redis.Pool {
	return t.pool
}

func (t *RedisPool) Close() {
	for i := 0; i < 10; i++ {
		if t.pool == nil {
			continue
		}

		err := t.pool.Close()
		if err == nil {
			break
		}
	}
}

func (t *RedisPool) Init(config RedisItemSchema, redisName string, seq int) {
	t.config = config
	t.redisName = redisName
	t.seq = seq
	t.Connect()
}

func (t *RedisPool) Connect() {
	if t.config.MaxReConnectionRetryTimes > 0 {
		t.connect(t.config.MaxReConnectionRetryTimes)
	} else {
		t.connect(maxRetryTimes)
	}
}

func (t *RedisPool) connect(retryTimes int) {

	if retryTimes <= 0 || t.pool != nil {
		return
	}

	config := t.config
	pool := &redis.Pool{
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		IdleTimeout: time.Duration(config.IdleTimeout) * time.Second,
		Wait:        config.Wait,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Host+":"+strconv.Itoa(config.Port))
			if err != nil {
				return nil, err
			}
			redis.DialDatabase(config.Database)
			if config.Password != "" {
				if _, err := c.Do("AUTH", config.Password); err != nil {
					_ = c.Close()
					return nil, err
				}
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

	key := "___redis_pool_test_flag___" + t.redisName
	//log.Println("key:=", key)
	redis1 := pool.Get()
	_, err := redis1.Do("SET", key, time.Now().String())
	if err != nil {
		toolSentry.Error(err)
	}

	_, err = redis1.Do("EXPIRE", key, maxExpireSeconds)
	if err != nil {
		toolSentry.Error(err)
	}

	for i := 0; i < 10; i++ {
		err := redis1.Close()
		if err == nil {
			break
		}
	}

	if err != nil {
		toolSentry.Error(err)
	}

	if err != nil {
		t.Log.Info(err)
		toolSentry.Error(err)
		t.connect(retryTimes - 1)
		return
	}

	//连接成功，且没有错误
	t.pool = pool
	return
}
