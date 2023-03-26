package toolMySQL

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	DBs      map[string]*gorm.DB = map[string]*gorm.DB{}
	dbsMutex sync.RWMutex
	LogHelp  *log.Helper
)

func DBsConnectOne(dbFlagName string, dsn string, config MySQLItemSchema) {

	if DBsGetOne(dbFlagName) == nil {
		DBsSetOne(dbFlagName, dsn, config)
	}
}

func DBsGetOne(dbFlagName string) *gorm.DB {

	dbsMutex.RLock()
	defer dbsMutex.RUnlock()

	dbList := _DBsGetList(dbFlagName)

	return dbList
}

func DBsPingOne(dbFlagName string) {

	db := DBsGetOne(dbFlagName)

	if db == nil {
		LogHelp.Error("db is not found")
		return
	}

	dbsMutex.Lock()
	defer dbsMutex.Unlock()

	if sqlDB, err := db.DB(); err == nil {
		if err := sqlDB.Ping(); err != nil {
			LogHelp.Errorf("err = %v\n", err)
		} else {
			LogHelp.Infof("DBsPingOne success. dbFlagName = %s", dbFlagName)
		}
	}
}

func _DBsGetList(dbFlagName string) *gorm.DB {

	_, ok := DBs[dbFlagName]
	if !ok {
		return nil
	}

	return DBs[dbFlagName]
}

func DBsSetOne(dbFlagName string, dsn string, config MySQLItemSchema) {
	dbsMutex.Lock()
	defer dbsMutex.Unlock()

	DBs[dbFlagName] = Connect(dsn, true, config, config.MaxRetryConnectTimes)
}

func Connect(dsn string, logMode bool, config MySQLItemSchema, retryTimes int) *gorm.DB {

	if retryTimes == 0 {
		return nil
	}

	var errorOccur error

	defer func() {
		if errorOccur != nil && retryTimes == 1 {
			LogHelp.Errorf("经过几次连接sql时，仍然失败，错误信息为：%v", errorOccur)
		}
	}()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		errorOccur = err
		return Connect(dsn, logMode, config, retryTimes-1)
	}

	if db.Error != nil {
		errorOccur = db.Error
		return Connect(dsn, logMode, config, retryTimes-1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		errorOccur = err
		return Connect(dsn, logMode, config, retryTimes-1)
	}

	// SetMaxIdleConns 设置空闲连接池中的最大连接数。
	sqlDB.SetMaxIdleConns(int(config.SetMaxIdleConns))

	// SetMaxOpenConns 设置数据库连接最大打开数。
	sqlDB.SetMaxOpenConns(int(config.SetMaxOpenConns))

	// SetConnMaxLifetime 设置可重用连接的最长时间
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.SetConnMaxLifetime))

	sqlDB.Ping()

	//db.LogMode(logMode)
	return db
}

func DBsCloseOne(dbFlagName string) {
	dbsMutex.Lock()
	defer dbsMutex.Unlock()

	if DBs[dbFlagName] != nil {
		sqlDB, _ := DBs[dbFlagName].DB()
		sqlDB.Close()
	}
}
