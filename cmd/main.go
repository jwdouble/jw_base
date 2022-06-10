package main

import (
	"context"
	"crypto/tls"
	_ "embed"
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"jw.lib/logx"

	pb "cmd/main.go/proto/generated_go"
	"cmd/main.go/service"
)

const addr = ":10000"

func main() {
	fmt.Printf("Starting server on %s\n", addr)
	tls, err := credentials.NewServerTLSFromFile("/static/tls.pem", "/static/tls.key")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(tls))
	//grpcServer := grpc.NewServer()
	pb.RegisterBaseServiceServer(grpcServer, service.NewBaseService())

	// 创建http路由和wg路由
	mux := http.NewServeMux()
	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("1234"))
	})

	gwmux := runtime.NewServeMux()

	dopts := []grpc.DialOption{grpc.WithTransportCredentials(tls)} // insecure.NewCredentials() 默认安全模式
	//dopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())} // insecure.NewCredentials() 默认安全模式
	err = pb.RegisterBaseServiceHandlerFromEndpoint(context.Background(), gwmux, addr, dopts)
	if err != nil {
		panic(err)
	}

	mux.Handle("/", gwmux)

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:      addr,
		Handler:   grpcHandleFunc(grpcServer, mux),
		TLSConfig: getTLSConfig(),
	}

	err = srv.Serve(conn)
	if err != nil {
		panic(err)
	}

	fmt.Println("exit")
}

// 选择合适的处理
func grpcHandleFunc(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有跨域请求
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func getTLSConfig() *tls.Config {
	pem, _ := os.ReadFile("/static/tls.pem")
	key, _ := os.ReadFile("/static/tls.key")
	logx.Info(string(pem))
	logx.Info(string(key))
	cert, err := tls.X509KeyPair(pem, key)
	if err != nil {
		logx.Error("tls load err: %s", err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{http2.NextProtoTLS}, // HTTP2 TLS支持
	}
}
