package main

import (
	"log"

	"github.com/gidyon/concurrency/gg/api"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewHelloServiceClient(conn)
	res, err := c.Echo(context.Background(), &api.HelloRequest{Name: "jack", Message: "happy new year folks"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", res.Reply)
}
