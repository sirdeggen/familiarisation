package rpc

import (
	"context"

	"github.com/sirdeggen/famailiarisation/tutorial2/internal/adapters/framework/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grcpa Adapter) GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grcpa.api.GetAddition(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected answer")
	}
	return &pb.Answer{
		Value: answer,
	}, nil
}

func (grcpa Adapter) GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grcpa.api.GetSubtraction(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected answer")
	}
	return &pb.Answer{
		Value: answer,
	}, nil
}

func (grcpa Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grcpa.api.GetMultiplication(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected answer")
	}
	return &pb.Answer{
		Value: answer,
	}, nil
}

func (grcpa Adapter) GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missing required")
	}

	answer, err := grcpa.api.GetDivision(req.A, req.B)
	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected answer")
	}
	return &pb.Answer{
		Value: answer,
	}, nil
}
