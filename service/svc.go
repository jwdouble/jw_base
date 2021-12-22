package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "cmd/main.go/proto/generated_go"
)

type MasterService struct {
	pb.UnimplementedMasterServiceServer
}

func (s *MasterService) Health(context.Context, *emptypb.Empty) (*pb.HealthMessage, error) {
	return &pb.HealthMessage{Version: "v1", Time: TimestampNow()}, nil
}
