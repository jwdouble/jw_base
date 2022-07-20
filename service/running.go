package service

import (
	"context"
	"fmt"
	"time"

	"jw.lib/conf"
	"jw.lib/timex"

	pb "cmd/main.go/proto/generated_go"
)

func (s *BaseService) GetRunningEnv(ctx context.Context, in *pb.GetRunningEnvReq) (*pb.GetRunningEnvResp, error) {
	var cfg = conf.EnvVar(in.Key)
	val := cfg.Value("empty")
	fmt.Println(time.Now().Format(timex.DateTimeFormat), val)
	return &pb.GetRunningEnvResp{Value: val}, nil
}
