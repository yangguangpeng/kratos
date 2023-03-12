package data

import (
	"helloworld/internal/conf"
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

	cleanup := mysqldb.Cleanup()

	return &Data{
		db: mysqldb,
	}, cleanup, nil
}
