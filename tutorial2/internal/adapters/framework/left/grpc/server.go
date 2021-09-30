package rpc

import (
	"log"
	"net"

	"github.com/sirdeggen/famailiarisation/tutorial2/internal/adapters/framework/left/grpc/pb"
	"github.com/sirdeggen/famailiarisation/tutorial2/internal/ports"
	"google.golang.org/grpc"
)

// Adapter is ...
type Adapter struct {
	api ports.APIPort
}

// NewAdapter is ...
func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

// Run is ...
func (grcpa Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	arithmeticServiceServer := grcpa
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000, %v", err)
	}
}
