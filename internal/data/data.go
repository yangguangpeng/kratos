package data

import (
	"helloworld/internal/conf"
	"helloworld/pkg/cache"
	"helloworld/pkg/db"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDemoRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db *db.Mysql
}

// NewData .
func NewData(bs *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger))

	mysqldb := db.New(db.WithConfig(bs), db.WithLog(l))
	redisService := cache.New(cache.WithConfig(bs), cache.WithLog(l))

	cleanup := func() {
		defer mysqldb.Close()
		defer redisService.Quit()
	}
	return &Data{
		db: mysqldb,
	}, cleanup, nil
}
