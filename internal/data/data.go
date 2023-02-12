package data

import (
	"gorm.io/gorm"
	"helloworld/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewDemoRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	db       *gorm.DB
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
	return &Data{
		db:db,
	}, cleanup, nil
}
