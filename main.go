package main

import (
	"test/handler"
	pb "test/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("test"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTestHandler(srv.Server(), new(handler.Test))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
