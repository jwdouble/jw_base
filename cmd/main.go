package main

import (
	pb "cmd/main.go/proto/generated_go"
	"cmd/main.go/service"
	"context"
	_ "embed"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"strings"
)

//go:embed generated_doc.swagger.json
var swaggerJSON string

const addr = "localhost:8081"

func main() {
	fmt.Printf("Starting server on %s\n", addr)
	grpcServer := grpc.NewServer()
	pb.RegisterBaseServiceServer(grpcServer, service.NewBaseService())

	// 创建http路由和wg路由
	mux := http.NewServeMux()
	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("1234"))
	})

	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterBaseServiceHandlerFromEndpoint(context.Background(), gwmux, addr, dopts)
	if err != nil {
		panic(err)
	}

	mux.Handle("/", gwmux)

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: grpcHandleFunc(grpcServer, mux),
	}

	err = srv.Serve(conn)
	if err != nil {
		panic(err)
	}

	fmt.Println("exit")
}

func grpcHandleFunc(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}
