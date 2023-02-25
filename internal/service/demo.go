package service

import (
	"context"
	"helloworld/internal/biz"

	pb "helloworld/api/helloworld/v1"
)

type DemoService struct {
	pb.UnimplementedDemoServer
	du *biz.DemoUsecase
}

func NewDemoService(DemoUsecase *biz.DemoUsecase) *DemoService {
	return &DemoService{du: DemoUsecase}
}

func (s *DemoService) CreateDemo(ctx context.Context, req *pb.CreateDemoRequest) (*pb.CreateDemoReply, error) {
	return &pb.CreateDemoReply{}, nil
}
func (s *DemoService) UpdateDemo(ctx context.Context, req *pb.UpdateDemoRequest) (*pb.UpdateDemoReply, error) {
	return &pb.UpdateDemoReply{}, nil
}
func (s *DemoService) DeleteDemo(ctx context.Context, req *pb.DeleteDemoRequest) (*pb.DeleteDemoReply, error) {
	return &pb.DeleteDemoReply{}, nil
}
func (s *DemoService) GetDemo(ctx context.Context, req *pb.GetDemoRequest) (*pb.GetDemoReply, error) {
	result, err := s.du.GetFormation(ctx, req.UserId)
	return &pb.GetDemoReply{Result: result + `80`}, err
}
func (s *DemoService) ListDemo(ctx context.Context, req *pb.ListDemoRequest) (*pb.ListDemoReply, error) {
	return &pb.ListDemoReply{}, nil
}
