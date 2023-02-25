package main

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/selector/p2c"
	"github.com/go-kratos/kratos/v2/selector/random"
	"github.com/go-kratos/kratos/v2/selector/wrr"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/hashicorp/consul/api"
	pb "helloworld/api/helloworld/v1"
)

func main() {

}

//服务发现与负载均衡
func balanceDemo() {
	selector.SetGlobalSelector(p2c.NewBuilder())
	selector.SetGlobalSelector(random.NewBuilder())
	selector.SetGlobalSelector(wrr.NewBuilder())

	// new consul client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	// new dis with consul client
	dis := consul.New(client)

	endpoint := "discovery:///myHello"
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(dis))
	if err != nil {
		panic(err)
	}
	demoClient := pb.NewDemoClient(conn)
	reply, err := demoClient.GetDemo(context.Background(), &pb.GetDemoRequest{UserId: 1})
	log.Infow(`reply`, reply, `err`, err)
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
