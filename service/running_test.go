package service

import (
	"testing"

	pb "cmd/main.go/proto/generated_go"
)

func TestBaseService_GetRunningEnv(t *testing.T) {
	r, err := baseSvc.GetRunningEnv(ctx, &pb.GetRunningEnvReq{Key: "COMMON_PASSWORD"})
	t.Log(r, err)
}
