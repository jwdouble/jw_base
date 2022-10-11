package main

import (
	"context"

	"jw.lib/conf"
	"jw.lib/rpcx"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"jw.lib/logx"

	pb "jw_base/proto/generated_go"
	"jw_base/service"
)

func main() {
	logx.KV("app", "jw-base")

	logx.Infof("server start, listen port %s", conf.SERVER_PORT.Value(":30001"))
	s := rpcx.New(conf.SERVER_PORT.Value(":30001"), func(grpcServer *grpc.Server, gwmux *runtime.ServeMux, dopts []grpc.DialOption) {
		pb.RegisterBaseServiceServer(grpcServer, service.NewBaseService())

		err := pb.RegisterBaseServiceHandlerFromEndpoint(context.Background(), gwmux, conf.SERVER_PORT.Value(":30001"), dopts)
		if err != nil {
			logx.Errorf(err, "pb.RegisterBaseServiceHandlerFromEndpoint")
		}
	})

	err := s.Run()
	if err != nil {
		logx.Errorf(err, "s.Run")
	}
}
