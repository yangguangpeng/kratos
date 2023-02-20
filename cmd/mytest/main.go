package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/selector/filter"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/hashicorp/consul/api"
	pb "helloworld/api/helloworld/v1"
)

func main() {

	// 创建路由 Filter：筛选版本号为"2.0.0"的实例
	filter := filter.Version("2.0.0")
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	// new dis with consul client
	dis := consul.New(client)
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///helloworld"),
		grpc.WithDiscovery(dis),
		// 通过 grpc.WithFilter 注入路由 Filter
		grpc.WithNodeFilter(filter),
	)
	fmt.Println(conn)
}

//golang的grpc例子与中间件jwt认证例子
func grpcAuthDemo() {
	testKey := `testKey`
	con, _ := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:8080"),
		grpc.WithMiddleware(
			jwt.Client(func(token *jwtv4.Token) (interface{}, error) {
				return []byte(testKey), nil
			}),
		),
	)
	demoClient := pb.NewDemoClient(con)
	reply, err := demoClient.GetDemo(context.Background(), &pb.GetDemoRequest{UserId: 2})
	log.Infow(`reply`, reply, `err`, err)
}
