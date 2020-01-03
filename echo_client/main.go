package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"

	pb "istio.io/troubleshoot/echo"
)

const (
	address = "localhost:443"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoClient(conn)

	name := "xxx"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Echo(ctx, &pb.EchoRequest{
		Name: name,
	})

	if err != nil {
		log.Fatalf("cannot send echo request: %v", err)
	}

	log.Printf("Echo: %s", r.Message)
}
