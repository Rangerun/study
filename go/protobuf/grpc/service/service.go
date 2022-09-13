package main

import (
	"context"
	pb "grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
} //服务对象

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello" + in.Nmae}, nil
}


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("err")
	}
	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Server(lis); err != nil {
		log.Fatal("err1")
	}

}