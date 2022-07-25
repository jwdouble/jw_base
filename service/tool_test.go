package service

import (
	"context"
	"testing"

	pb "cmd/main.go/proto/generated_go"
)

func TestBaseService_ParseJwt(t *testing.T) {
	r, err := baseSvc.ParseJwt(context.Background(), &pb.ParseJwtReq{Jwt: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZW4iOiJqd3Rlc3QiLCJhdWQiOiIqIiwic3ViIjoiU1lTVEVNIiwianRpIjoiWDdaRzNYMDE0TEMiLCJpYXQiOi0xLCJleHQiOnsicHJpdmlsZWdlIjp0cnVlfSwicHJqIjoiand0ZXN0IiwiaXNzIjoiKiIsInVpZCI6IlNZU1RFTSJ9.lBm2zX7g3pR1hJbuRpqyU0-PFvrmfs9MEGCtYGY1RjI"})
	t.Log(r, err)
}
