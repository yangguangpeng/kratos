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
	options   *options
	toolMysql *toolMySQL.InitMySQL
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
	m := &Mysql{options: &o}
	m.initMysql()
	return m
}

func (m *Mysql) initMysql() {
	fmt.Println(`Mysql initMysql...`)
	o := m.options
	//config := m.options.config
	//songguoMaster := config.GetMysql().Songguo.GetMaster()
	mysqlConfig := make(map[string]toolMySQL.MySQLItemSchema)
	mysqlConfig[SONGGUO_MASTER] = toolMySQL.MySQLItemSchema{
		Dsn:                  m.makeSongguoMasterDSN(),
		MaxRetryConnectTimes: 3,
		//SetMaxIdleConns:      songguoMaster.GetSetMaxIdleConns(),
		//SetMaxOpenConns:      songguoMaster.GetSetMaxOpenConns(),
		//SetConnMaxLifetime:   songguoMaster.GetSetConnMaxLifetime(),
	}

	toolMysql := &toolMySQL.InitMySQL{
		MySQLInfo: mysqlConfig,
		Log:       o.log}
	fmt.Println(`Mysql initMysql 2222...`)
	toolMysql.Init()
	m.toolMysql = toolMysql
}

func (m *Mysql) makeSongguoMasterDSN() string {
	o := m.options

	//return fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s`, `root`,
	//	`admin123`, `127.0.0.1`, 3306, `test`)
	songguoMysql := o.config.GetMysql().GetSongguo()
	songguoMaster := songguoMysql.GetMaster()

	username := songguoMaster.GetUsername()
	password := songguoMaster.GetPassword()
	host := songguoMaster.GetHost()
	port := songguoMaster.GetPort()
	dbname := songguoMysql.GetDbName()
	if username == `` || password == `` || host == `` || port == 0 || dbname == `` {
		m.options.log.Error(`读取配置songguoMaster失败`)
		return ``
	}
	return fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s`, username,
		password,
		host,
		port,
		dbname)
}

func (m *Mysql) GetSongguoMaster() *gorm.DB {
	gdb, ok := toolMySQL.DBs[SONGGUO_MASTER]
	if !ok {
		m.options.log.Errorf(`对象%s不存在`, SONGGUO_MASTER)
		return nil
	}
	return gdb
}

func (m *Mysql) GetSlave() {

}

func (m *Mysql) Close() {
	m.toolMysql.Close()
}
