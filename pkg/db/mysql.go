package db

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"helloworld/internal/conf"
	"helloworld/pkg/db/toolMySQL"
)

const SONGGUO_MASTER = `songguoMaster`

type Mysql struct {
	log *log.Helper
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

func New(opts ...Option) *Mysql {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}
	initMysql(o)
	return &Mysql{o.log}
}

func initMysql(o options) {
	config := make(map[string]toolMySQL.MySQLItemSchema)

	config[SONGGUO_MASTER] = toolMySQL.MySQLItemSchema{
		Dsn:                  makeSongguoMasterDSN(o),
		MaxRetryConnectTimes: 3,
	}

	mysql := &toolMySQL.InitMySQL{
		MySQLInfo: config,
		Log:       o.log}

	mysql.Init()

}

func makeSongguoMasterDSN(o options) string {
	songguoMysql := o.config.GetMysql().Songguo
	songguoMaster := songguoMysql.GetMaster()
	return fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`, songguoMaster.GetUsername(),
		songguoMaster.GetPassword(),
		songguoMaster.GetHost(),
		songguoMaster.GetPort(),
		songguoMysql.GetDbName())
}

func (m *Mysql) GetSongguoMaster() *gorm.DB {
	gdb, ok := toolMySQL.DBs[SONGGUO_MASTER]
	if !ok {
		m.log.Errorf(`对象%s不存在`, SONGGUO_MASTER)
		return nil
	}
	return gdb
}

func (m *Mysql) GetSlave() {

}

func Cleanup(toolMysql *toolMySQL.InitMySQL) func() {
	return func() {
		toolMysql.Close()
	}
}
