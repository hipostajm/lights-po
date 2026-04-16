package main

import (
	"lightswitch-service/internal/adapters/brodcast"
	"lightswitch-service/internal/adapters/repository"
	"lightswitch-service/internal/adapters/server"
	"lightswitch-service/internal/core/services"
	"log"
	"net"
	"sync"
	"time"

	lightswitchv1 "proto/lightswitch/v1"

	"google.golang.org/grpc"
)

func main(){
	port := "6741"

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil{
		log.Fatal(err)
	}
	
	grpcServer := grpc.NewServer()

	repo := repository.NewLightSwitchInMemoryRepository()
	brodacst := brodcast.NewLiLightSwitchBrodcastImpl(sync.Mutex{}, "tcp://nats:1883", "lightswitch-broadcaster")
	service := services.NewLightSwithcService(repo, brodacst, time.Minute)
	server := server.NewLightSwitchServer(service)
	
	lightswitchv1.RegisterLightswitchServiceServer(grpcServer, server)

	log.Println("running on port:", port)
	err = grpcServer.Serve(lis)
	if err != nil{
		log.Fatal(err)
	}
}
