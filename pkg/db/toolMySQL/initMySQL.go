package toolMySQL

import (
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

var (
	maxHealthCheckDuration = 3 * 1000 //20秒
)

type MySQLItemSchema struct {
	Dsn                  string
	SetMaxIdleConns      int32
	SetMaxOpenConns      int32
	SetConnMaxLifetime   int64
	MaxRetryConnectTimes int32
}

type MySQLSchema map[string]MySQLItemSchema

type InitMySQL struct {
	MySQLInfo          MySQLSchema
	systemQuit         chan struct{}
	systemQuitFinished chan struct{}
	Log                *log.Helper
}

func (t *InitMySQL) Init() {
	t.systemQuit = make(chan struct{}, 1)
	t.systemQuitFinished = make(chan struct{}, 1)
	t.connect()
	go t.HealthCheck(maxHealthCheckDuration)
}

func (t *InitMySQL) connect() {

	for dbFlagName, schema := range t.MySQLInfo {
		DBsConnectOne(dbFlagName, schema.Dsn, schema)
		DBsPingOne(dbFlagName)
	}

}

// HealthCheck 健康检查 这里的健康检查只是简单的ping检查，connect()本身会对启动时连接失败的场景每十秒做一次是否需要初始化的检查。
// 目前还没有思考怎么减少断开连接后更实时的连接。主要是因为连接每十秒一次。
func (t *InitMySQL) HealthCheck(healthCheckDuration int) {
	t.connect()
	//定时检查
	select {
	case <-t.systemQuit:
		log.Info("InitMySQL.healthCheck exited")
		t.systemQuitFinished <- struct{}{}
		return
	case <-time.After(time.Duration(healthCheckDuration) * time.Millisecond):
		t.HealthCheck(healthCheckDuration)
	}
}

func (t *InitMySQL) Close() {

	for dbFlagName, _ := range t.MySQLInfo {
		DBsCloseOne(dbFlagName)
	}
	t.systemQuit <- struct{}{}
	log.Info("InitMySQL.Close exited")
	<-t.systemQuitFinished
}
