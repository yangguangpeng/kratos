package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Demo struct {
	UserID int64
	Age int8
}

type DemoRepo interface {
	FindByID(context.Context, int64) (*Demo, error)
}


type DemoUsecase struct {
	repo DemoRepo
	log  *log.Helper
}

func NewDemoUsecase(repo DemoRepo, logger log.Logger) *DemoUsecase {
	return &DemoUsecase{repo: repo, log: log.NewHelper(logger)}
}

//实现业务逻辑的层
func (du *DemoUsecase) GetFormation(ctx context.Context, userID int64) (string, error) {
	demo,err := du.repo.FindByID(ctx, userID)
	ret := `userID:`+string(userID)+`,年龄：`+ string(demo.Age)
	return ret, err
}