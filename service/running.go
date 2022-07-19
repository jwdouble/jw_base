package service

import (
	"context"

	"jw.lib/conf"

	pb "cmd/main.go/proto/generated_go"
)

func (s *BaseService) GetRunningEnv(ctx context.Context, in *pb.GetRunningEnvReq) (*pb.GetRunningEnvResp, error) {
	var cfg = conf.EnvVar(in.Key)
	val := cfg.Value("empty")
	return &pb.GetRunningEnvResp{Value: val}, nil
}
