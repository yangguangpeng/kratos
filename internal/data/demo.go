package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
	"helloworld/internal/data/model"
)

type demoRepo struct {
	data *Data
	log  *log.Helper
}


func (d *demoRepo) FindByID(context.Context, int64) (*biz.Demo, error) {
	userInfo, err := model.UsersMgr(d.data.db).FetchByPrimaryKey(1)
	return &biz.Demo{
		Age: userInfo.Age,
	}, err
}

// NewDemoRepo .
func NewDemoRepo(data *Data, logger log.Logger) biz.DemoRepo {
	return &demoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
