package service

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "cmd/main.go/proto/generated_go"
)

type BaseService struct {
	pb.UnimplementedBaseServiceServer
}

func NewBaseService() *BaseService {
	return &BaseService{}
}

func (s *BaseService) Health(context.Context, *emptypb.Empty) (*pb.HealthMessage, error) {
	return &pb.HealthMessage{Version: "v1", Time: TimestampNow()}, nil
}

func TimestampNow() *pb.Timestamp {
	current := time.Now()
	return &pb.Timestamp{Seconds: current.Unix(), Nanos: int32(current.Nanosecond())}
}
