package main

import (
	"context"
	_ "embed"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	pb "cmd/main.go/proto/generated_go"
	"cmd/main.go/service"
)

//go:embed generated_doc.swagger.json
var swaggerJSON string

func main() {
	log.Println("jw-base start run")
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	http.

}
