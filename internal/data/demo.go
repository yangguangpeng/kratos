package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
	model2 "helloworld/internal/data/repositories/mysql/model"
	"strconv"
)

type demoRepo struct {
	data *Data
	log  *log.Helper
}

func (d *demoRepo) FindByID(c context.Context, userID int64) (*biz.Demo, error) {
	d.log.Info(`请求参数` + strconv.Itoa(int(userID)))
	odb := d.data.db.GetSongguoMaster()
	userInfo, _ := model2.UsersMgr(odb).FetchByPrimaryKey(uint32(userID))
	retAge := userInfo.Age
	////规范处理，写个拼接名称
	//d.data.redisCache.KeyName = `userID:`+strconv.Itoa(int(userID))
	//age,err := redis.Int(d.data.redisCache.Get())
	//
	//var retAge int8
	//retAge = 0
	//
	//if err == nil {
	//	retAge = int8(age)
	//} else {
	//	userInfo, _ := model.UsersMgr(d.data.db).FetchByPrimaryKey(uint32(userID))
	//	if userInfo.Age == 0 {
	//		return &biz.Demo{}, v1.ErrorUserNotFound(`user %s not found`,strconv.Itoa(int(userID)))
	//	}
	//	retAge = userInfo.Age
	//}

	return &biz.Demo{
		Age: retAge,
	}, nil
}

// NewDemoRepo .
func NewDemoRepo(data *Data, logger log.Logger) biz.DemoRepo {
	return &demoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
