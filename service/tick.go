package service

import (
	"log"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cmd/main.go/proto/generated_go"
)

func (s *MasterService) Tick(_ *emptypb.Empty, stream pb.MasterService_TickServer) error {
	t := &pb.Tick{
		Tick: TimestampNow(),
	}

	select {
	case <-time.Tick(time.Minute):
		break
	case <-time.Tick(time.Second):
		err := stream.Send(t)
		log.Fatal(err)
	}

	return nil
}

func TimestampNow() *timestamppb.Timestamp {
	t := time.Now()
	return &timestamppb.Timestamp{Seconds: int64(t.Second()), Nanos: int32(t.Nanosecond())}
}
