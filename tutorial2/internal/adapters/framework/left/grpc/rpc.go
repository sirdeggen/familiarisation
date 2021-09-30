package rpc

import (
	"context"
	"log"

	"github.com/sirdeggen/famailiarisation/tutorial2/internal/adapters/framework/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetAddition is ...
func (grcpa Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grcpa.api.GetAddition(req.A, req.B)
	log.Println("MOOSE", answer, err, req)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected answer")
	}
	return &pb.Answer{
		Value: answer,
	}, nil
}

// GetSubtraction is ...
func (grcpa Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grcpa.api.GetSubtraction(req.A, req.B)
	log.Println("MOOSE", answer, err, req)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected answer")
	}
	return &pb.Answer{
		Value: answer,
	}, nil
}

// GetMultiplication is ...
func (grcpa Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grcpa.api.GetMultiplication(req.A, req.B)
	log.Println("MOOSE", answer, err, req)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected answer")
	}
	return &pb.Answer{
		Value: answer,
	}, nil
}

// GetDivision is ...
func (grcpa Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grcpa.api.GetDivision(req.A, req.B)
	log.Println("MOOSE", answer, err, req)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected answer")
	}
	return &pb.Answer{
		Value: answer,
	}, nil
}
