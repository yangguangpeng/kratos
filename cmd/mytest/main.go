package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	pb "helloworld/api/helloworld/v1"
)

func main() {
	grpcAuthDemo()
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
