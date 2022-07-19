package service

import (
	"fmt"
	"testing"

	pb "cmd/main.go/proto/generated_go"
)

func TestBaseService_GetRunningEnv(t *testing.T) {
	r, err := baseSvc.GetRunningEnv(ctx, &pb.GetRunningEnvReq{Key: "COMMON_PASSWORD"})
	fmt.Println(r, "\r\n", err)
}
