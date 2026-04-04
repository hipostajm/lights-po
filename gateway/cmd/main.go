package main

import (
	"gateway/internal/handler"
	"gateway/internal/repository"
	"gateway/internal/service"
	"gateway/internal/utils"
	"log"
	"time"

	pb "proto/lightswitch/v1"

	"github.com/gofiber/fiber/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main(){
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(logger.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	grpcClient, err := grpc.NewClient("lightswitch-service:6741", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatal(err)
	}
	defer grpcClient.Close()

	lightSwitchGrpcClient := pb.NewLightswitchServiceClient(grpcClient)	
	lightsRepository := repository.NewLightsRepositoryGrpc(lightSwitchGrpcClient, time.Second)
	lightsService := service.NewLightsService(lightsRepository)
	lightsHandler := handler.NewLightsHandler(lightsService)

	lights := v1.Group("/lights")

	lights.Post("/", lightsHandler.AddLightSwitch)
	lights.Patch("/:id", lightsHandler.ToggleSwitchHandler)
	// lights.Get("/:id/status", lightsHandler.GetLightSwitchStats)

	log.Fatal(app.Listen(":8080"))
}
