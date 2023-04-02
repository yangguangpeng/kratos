package cache

import (
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/conf"
	"helloworld/pkg/cache/toolRedis"
)

const BIG_CACHE_MASTER = `bigCacheMaster`

type RedisService struct {
	options   *options
	redisBase *toolRedis.InitRedis
}

type Option func(*options)

type options struct {
	config *conf.Bootstrap
	log    *log.Helper
}

func WithConfig(config *conf.Bootstrap) Option {
	return func(opts *options) {
		opts.config = config
	}
}

func WithLog(log *log.Helper) Option {
	return func(opts *options) {
		opts.log = log
	}
}

func New(opts ...Option) *RedisService {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}
	r := &RedisService{options: &o}
	r.initRedis()
	return r
}

func (r *RedisService) initRedis() {

	redisInfo := toolRedis.RedisSchema{}
	var schemaSlice []toolRedis.RedisItemSchema
	bigCacheMaster := r.options.config.GetRedis().GetBigcache().GetMaster()
	bigCacheMasterHost := bigCacheMaster.GetHost()
	bigCacheMasterPort := bigCacheMaster.GetPort()
	if bigCacheMasterHost == `` {
		r.options.log.Error(`读取配置bigCacheMaster HOST失败`)
		return
	}
	if bigCacheMasterPort == 0 {
		r.options.log.Error(`读取配置bigCacheMaster PORT失败`)
		return
	}
	r.options.log.Info(`bigCacheMaster.GetHost()`, bigCacheMaster.GetHost())
	schemaSlice = append(schemaSlice, toolRedis.RedisItemSchema{
		Host: `127.0.0.1`,
		Port: 6379,
	})
	redisInfo[BIG_CACHE_MASTER] = schemaSlice
	// redis连接池中间件，仅单节点
	r.redisBase = &toolRedis.InitRedis{
		RedisInfo: redisInfo,
		Log:       r.options.log,
	}
	r.redisBase.Init()
}

func (r *RedisService) Quit() {
	r.redisBase.Quit()
}
