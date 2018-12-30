package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gidyon/concurrency/gg/api"

	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	var hs *helloServer
	api.RegisterHelloServiceServer(srv, hs)

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("could not listen on tcp :9090: %v\n", err)
	}
	log.Println("Server started on port :9090, waiting for rpc calls...")
	log.Fatalln(srv.Serve(l))
}

type helloServer struct{}

func (hs *helloServer) Echo(ctx context.Context, req *api.HelloRequest) (*api.HelloResponse, error) {
	helloResponse := &api.HelloResponse{
		Reply: fmt.Sprintf("server got msg: %q, from: %q", req.Message, req.Name),
	}
	return helloResponse, nil
}
