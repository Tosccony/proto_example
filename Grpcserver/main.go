package main

import (
	"context"
	"log"
	"net"

	pb "github.com/tosccony/proto_example"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Create a struct and embed our UnimplementedExampleServer
// We provide a full implementation to the methods that this embedded struct specifies down below
type server struct {
	pb.UnimplementedExampleServer
}

func (s *server) GetMenu(*pb.MenuRequest, pb.Example_GetMenuServer) error {
	return status.Error(codes.Unimplemented, "method GetMenu not implemented")
}

func (s *server) PlaceOrder(context.Context, *pb.Order) (*pb.Receipt, error) {
	return nil, status.Error(codes.Unimplemented, "method PlaceOrder not implemented")
}

func (s *server) GetOrderStatus(context.Context, *pb.Receipt) (*pb.OrderStatus, error) {
	return nil, status.Error(codes.Unimplemented, "method GetOrderStatus not implemented")
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExampleServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
