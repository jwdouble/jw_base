package service

import (
	pb "cmd/main.go/proto/generated_go"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"google.golang.org/protobuf/types/known/structpb"
)

func (s *BaseService) ParseJwt(ctx context.Context, req *pb.ParseJwtReq) (*pb.ParseJwtResp, error) {
	var hmacSampleSecret []byte
	token, err := jwt.Parse(req.Jwt, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	m := token.Claims.(jwt.MapClaims)

	for k, v := range m {
		fmt.Println(k, v)
	}

	c, err := structpb.NewStruct(m)
	if err != nil {
		return nil, err
	}

	return &pb.ParseJwtResp{Claims: c}, nil
}
