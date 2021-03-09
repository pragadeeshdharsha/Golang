package main

import (
	"context"
	pb "grpctest/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedCalcServiceServer
}

func main() {

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterCalcServiceServer(srv, &server{})
	//for serialising and de-serialinsing data
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *pb.ClientRequest) (*pb.ServerResponse, error) {

	a, b := request.GetFirstNum(), request.GetSecondNum()
	result := a + b

	return &pb.ServerResponse{Result: result}, nil

}

func (s *server) Multiply(ctx context.Context, request *pb.ClientRequest) (*pb.ServerResponse, error) {

	a, b := request.GetFirstNum(), request.GetSecondNum()
	result := a * b

	return &pb.ServerResponse{Result: result}, nil
}
