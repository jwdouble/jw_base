package main

import (
	pb "cmd/main.go/proto/generated_go"
	"cmd/main.go/service"
	"context"
	"crypto/tls"
	_ "embed"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"net"
	"net/http"
	_ "net/http/pprof"
	"strings"
)

//go:embed generated_doc.swagger.json
var swaggerJSON string

const addr = ":10000"

func main() {
	fmt.Printf("Starting server on %s\n", addr)
	//tls, err := credentials.NewServerTLSFromFile("/static/tls.pem", "/static/tls.key")
	//if err != nil {
	//	panic(err)
	//}

	//grpcServer := grpc.NewServer(grpc.Creds(tls))
	grpcServer := grpc.NewServer()
	pb.RegisterBaseServiceServer(grpcServer, service.NewBaseService())

	// 创建http路由和wg路由
	mux := http.NewServeMux()
	mux.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("1234"))
	})

	gwmux := runtime.NewServeMux()

	//dopts := []grpc.DialOption{grpc.WithTransportCredentials(tls)} // insecure.NewCredentials() 默认安全模式
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())} // insecure.NewCredentials() 默认安全模式
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
		//TLSConfig: getTLSConfig(),
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func getTLSConfig() *tls.Config {
	cert, _ := ioutil.ReadFile("/static/tls.pem")
	key, _ := ioutil.ReadFile("/static/tls.key")
	var demoKeyPair *tls.Certificate
	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		grpclog.Fatalf("TLS KeyPair err: %v\n", err)
	}
	demoKeyPair = &pair
	return &tls.Config{
		Certificates: []tls.Certificate{*demoKeyPair},
		NextProtos:   []string{http2.NextProtoTLS}, // HTTP2 TLS支持
	}
}
