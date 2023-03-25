package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/gomodule/redigo/redis"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/data/repositories/redis/bigCache"
	"strconv"
)

type Demo struct {
	UserID int64
	Age    int8
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
	du.log.WithContext(ctx).Info(`userid:` + strconv.Itoa(int(userID)))
	resStr, _ := redis.String(bigCache.UserInstance().UseSlave().SetKeyName(userID).Get())
	du.log.WithContext(ctx).Info(`res:` + resStr)
	return resStr, nil
	//traceId := `init`
	//if header, ok := transport.FromServerContext(ctx); ok {
	//	traceId = header.RequestHeader().Get(`Mytraceid`)
	//}
	//logger := log.With(log.NewStdLogger(os.Stdout),
	//	"ts", log.DefaultTimestamp,
	//	"defaultCaller", log.DefaultCaller,
	//	"caller", log.Caller,
	//	"service.id", 1,
	//	"service.name", 1,
	//	"service.version", 1,
	//	"trace.id", traceId,
	//	"span.id", tracing.SpanID(),
	//)
	//du.log = log.NewHelper(logger)
	if header, ok := transport.FromServerContext(ctx); ok {
		du.log.Infow(`Mytraceid`, header.RequestHeader().Get(`Mytraceid`))
		du.log.WithContext(ctx).Info(`come in`)
	}

	du.log.WithContext(ctx).Info(`userID:这里:` + strconv.Itoa(int(userID)))

	return ``, nil
	demo, err := du.repo.FindByID(ctx, userID)
	if v1.IsUserNotFound(err) {
		//取出错误的具体信息
		return errors.FromError(err).GetMessage(), nil
	}
	ret := `userID:` + strconv.Itoa(int(userID)) + `,年龄：` + strconv.Itoa(int(demo.Age))
	return ret, err
}
