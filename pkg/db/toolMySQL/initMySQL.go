package toolMySQL

import (
	"time"
)

var (
	maxHealthCheckDuration = 20 * 1000 //20秒
)

type MySQLItemSchema struct {
	Dsn                  string
	SetMaxIdleConns      int
	SetMaxOpenConns      int
	SetConnMaxLifetime   int64
	MaxRetryConnectTimes int
}

type MySQLSchema map[string]MySQLItemSchema

type InitMySQL struct {
	MySQLInfo []MySQLSchema
}

func (t *InitMySQL) Init() {
	t.connect()
	go t.HealthCheck(maxHealthCheckDuration)
}

func (t *InitMySQL) connect() {

	for _, item := range t.MySQLInfo {

		for dbFlagName, schema := range item {
			DbsInit(dbFlagName)

			DBsConnectOne(dbFlagName, schema.Dsn, schema)
			DBsPingOne(dbFlagName)
		}

	}

}

// HealthCheck 健康检查 这里的健康检查只是简单的ping检查，connect()本身会对启动时连接失败的场景每十秒做一次是否需要初始化的检查。
// 目前还没有思考怎么减少断开连接后更实时的连接。主要是因为连接每十秒一次。
func (t *InitMySQL) HealthCheck(healthCheckDuration int) {
	t.connect()
	//定时检查
	select {
	case <-time.After(time.Duration(healthCheckDuration) * time.Millisecond):
		t.HealthCheck(healthCheckDuration)
	}
}
