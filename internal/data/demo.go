package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomodule/redigo/redis"
	"helloworld/internal/biz"
	"helloworld/internal/data/model"
	"strconv"
)

type demoRepo struct {
	data *Data
	log  *log.Helper
}


func (d *demoRepo) FindByID(c context.Context, userID int64) (*biz.Demo, error) {
	//规范处理，写个拼接名称
	d.data.redisCache.KeyName = `userID:`+strconv.Itoa(int(userID))
	age,err := redis.Int(d.data.redisCache.Get())

	fmt.Println(age,"打印redis结果")
	var retAge int8
	retAge = 0

	if err == nil {
		retAge = int8(age)
	}else {
		userInfo, _ := model.UsersMgr(d.data.db).FetchByPrimaryKey(uint32(userID))
		retAge = userInfo.Age
	}

	return &biz.Demo{
		Age: retAge,
	}, err
}

// NewDemoRepo .
func NewDemoRepo(data *Data, logger log.Logger) biz.DemoRepo {
	return &demoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
