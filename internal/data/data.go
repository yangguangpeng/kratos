package data

import (
	"gorm.io/gorm"
	"helloworld/internal/conf"
	"helloworld/internal/data/redis"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDemoRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db       *gorm.DB
	redisCache    *redis.Base
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "helloWorld-service/data"))

	db, err := NewGorm(c.Database.Source)

	cleanup := func() {
		sqlDB, err := db.DB()
		if err != nil {
			l.Error(err)
		} else {
			if err = sqlDB.Close(); err != nil {
				l.Error(err)
			}
		}
		log.NewHelper(logger).Info("closing the data resources")
	}

	if err != nil {
		l.Errorf("gorm err %v ", err)
		return nil, nil, err
	}

	redisCache, err1 := redis.NewBase(c.Redis)

	if err != nil {
		l.Errorf("redis err %v ", err1)
		return nil, nil, err1
	}

	return &Data{
		db:db,
		redisCache:redisCache,
	}, cleanup, nil
}
