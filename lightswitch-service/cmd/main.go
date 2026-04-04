package main

import (
	"lightswitch-service/internal/adapters/repository"
	"lightswitch-service/internal/adapters/server"
	"lightswitch-service/internal/core/services"
	"log"
	"net"

	lightswitchv1 "proto/lightswitch/v1"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main(){
	
	nc, err := nats.Connect("nats://nats:4222")

	if err != nil{
		log.Fatal(err)
	}
	defer nc.Drain()

	port := "6741"

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil{
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	repo := repository.NewLightSwitchInMemoryRepository()
	service := services.NewLightSwithcService(repo)
	server := server.NewLightSwitchServer(service, nc)
	
	lightswitchv1.RegisterLightswitchServiceServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil{
		log.Fatal(err)
	}
	log.Println("running on port:", port)
}
