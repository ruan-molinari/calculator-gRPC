package main

import (
	"context"
	"log"
	"net"
	"regexp"
	"strings"

	"github.com/ruan-molinari/calculator-gRPC/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, in *pb.Input) (*pb.Output, error) {
	return &pb.Output{
		Result: in.Operand1 + in.Operand2,
	}, nil
}

func (s *server) Divide(ctx context.Context, in *pb.Input) (*pb.Output, error) {
	if in.Operand2 == 0 {
		return nil, status.Error(
			codes.InvalidArgument, "Cannot divide by zero",
		)
	}

	return &pb.Output{
		Result: in.Operand1 / in.Operand2,
	}, nil
}

func (s *server) Sum(ctx context.Context, in *pb.RepeatedInput) (*pb.Output, error) {
	var sum int64
	
	for _, num := range in.Numbers {
		sum += num
	}
	return &pb.Output{
		Result: sum,
	}, nil
} 

func (s *server) Multiply(ctx context.Context, in *pb.Input) (*pb.Output, error) {
	return &pb.Output{
		Result: in.Operand1 * in.Operand2,
	}, nil
}

func (s *server) Subtract(ctx context.Context, in *pb.Input) (*pb.Output, error) {
	return &pb.Output{
		Result: in.Operand1 - in.Operand2,
	}, nil
}

func main() {
	// Instanciates a new listener
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln("failed to create listener", err)
	}
	
	// Instanciates a new server
	s := grpc.NewServer()
	// Enable reflection for testing and debugging
	reflection.Register(s)
	
	// Registers the calculator server
	pb.RegisterCalculatorServer(s, &server{})
	// Tries to serve the calculator on the listener
	if err := s.Serve(listener); err != nil {
		log.Println("faled to serve: ", err)
	}
	log.Println("server running at port 8080")
}
