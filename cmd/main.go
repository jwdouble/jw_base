package main

import (
	_ "embed"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "cmd/main.go/proto/generated_go"
	"cmd/main.go/service"
)

//go:embed generated_doc.swagger.json
var swaggerJSON string

func main() {
	log.Println("jw-base start run")
	lis, err := net.Listen("tcp", "localhost:9876")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMasterServiceServer(s, &service.MasterService{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatal("server err:", err)
	}
}
