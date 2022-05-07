package main

import (
	"log"
	"net"
	"os"
	"panagiotisptr/authentication/protos"
	"panagiotisptr/authentication/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := log.New(os.Stdout, "messaging-service", log.Lshortfile)

	gs := grpc.NewServer()
	cs, err := server.NewAuthenticationServer(logger)
	if err != nil {
		panic(err)
	}

	protos.RegisterAuthenticationServer(gs, cs)

	reflection.Register(gs)

	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Fatalf("Unable to listen. Error: %v", err)
		os.Exit(1)
	}

	gs.Serve(list)
}
