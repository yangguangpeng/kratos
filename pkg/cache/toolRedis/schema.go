package toolRedis

type RedisItemSchema struct {
	Host                      string
	Port                      string
	Password                  string
	Database                  int
	MaxIdle                   int
	MaxActive                 int
	IdleTimeout               int
	Wait                      bool
	MaxReConnectionRetryTimes int
}

type RedisSchema map[string][]RedisItemSchema
