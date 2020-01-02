//go:generate protoc -I ../echo --go_out=plugins=grpc:../echo ../echo/echo.proto
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "istio.io/troubleshoot/echo"
)

const (
	port = ":6666"
)

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.EchoReply{Message: "Echo " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
