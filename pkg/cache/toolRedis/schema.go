package toolRedis

type RedisItemSchema struct {
	Host                      string
	Port                      int
	Password                  string
	Database                  int
	MaxIdle                   int
	MaxActive                 int
	IdleTimeout               int
	Wait                      bool
	MaxReConnectionRetryTimes int
}

type RedisSchema map[string][]RedisItemSchema
