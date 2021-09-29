package ports

import (
	"context"

	pb "github.com/sirdeggen/famailiarisation/tutorial2/internal/adapters/framework/left/grpc/pb"
)

type GRCPPort interface {
	Run()
	GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
}
