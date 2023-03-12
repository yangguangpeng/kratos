package main

import (
	"fmt"
	"helloworld/internal/data/model"
	"helloworld/pkg/db/toolMySQL"
	"time"
)

func main() {

	config := make(map[string]toolMySQL.MySQLItemSchema)
	config[`songguo`] = toolMySQL.MySQLItemSchema{
		Dsn:                  `root:admin123@tcp(127.0.0.1:3306)/test`,
		MaxRetryConnectTimes: 3,
	}
	mysqlInfo := []toolMySQL.MySQLSchema{}

	mysqlInfo = append(mysqlInfo, config)

	mysql := &toolMySQL.InitMySQL{
		mysqlInfo}

	mysql.Init()
	db := toolMySQL.DBs["songguo"]

	userID := 1
	userInfo, _ := model.UsersMgr(db).FetchByPrimaryKey(uint32(userID))
	retAge := userInfo.Age
	fmt.Println(retAge)

	time.Sleep(time.Duration(2000) * time.Second)
}
