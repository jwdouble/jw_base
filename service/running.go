package service

import (
	"context"
	"time"

	"jw.lib/conf"
	"jw.lib/logx"
	"jw.lib/timex"

	pb "jw_base/proto/generated_go"
)

func (s *BaseService) GetRunningEnv(ctx context.Context, in *pb.GetRunningEnvReq) (*pb.GetRunningEnvResp, error) {
	var cfg = conf.EnvVar(in.Key)
	val := cfg.Value("empty")
	logx.Debugf(time.Now().Format(timex.DateTimeFormat), val)
	return &pb.GetRunningEnvResp{Value: val}, nil
}
