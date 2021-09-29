package main

import (
	"log"
	"os"

	"github.com/sirdeggen/famailiarisation/tutorial2/internal/adapters/app/api"
	"github.com/sirdeggen/famailiarisation/tutorial2/internal/adapters/core/arithmetic"
	gRPC "github.com/sirdeggen/famailiarisation/tutorial2/internal/adapters/framework/left/grpc"
	"github.com/sirdeggen/famailiarisation/tutorial2/internal/adapters/framework/right/db"
	"github.com/sirdeggen/famailiarisation/tutorial2/internal/ports"
)

func main() {
	var err error

	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRCPPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsName := os.Getenv("DS_NAME")
	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsName)
	if err != nil {
		log.Fatalf("Failed to init dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()

}
