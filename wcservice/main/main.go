package main

import (
	"fmt"
	"wcservice/api"
	"wcservice/configuration"
	"wcservice/grpc"
	"log"
	"net"
)

func main() {

	cfg, err := configuration.ReadConfigFromEnv()
	if err != nil {
		log.Fatalf("failed to read env: %v", err)
	}

	var errs chan error

	r := api.Setup()

	go func() {
		errs <- r.Run(fmt.Sprintf(":%d", cfg.RestPort))
	}()

	lis, err := net.Listen("",fmt.Sprintf(":%d", cfg.GRPCPort))
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}

	s := grpc.Setup()

	go func() {
		fmt.Printf("\ngRPC server Listening at %v\n", lis.Addr())
		errs <- s.Serve(lis)
	}()

	if err = <-errs; err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	s.GracefulStop()
}
