package main

import (
	"gateway/internal/handler"
	"gateway/internal/repository"
	"gateway/internal/service"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main(){
	app := fiber.New()
	api := app.Group("/api")
	v1 := api.Group("/v1")

	lightsRepository := repository.NewLightsRepositoryInMemory()
	lightsService := service.NewLightsService(lightsRepository)
	lightsHandler := handler.NewLightsHandler(lightsService)

	lights := v1.Group("/lights")

	lights.Post("/", lightsHandler.AddLightSwitch)
	lights.Patch("/:id", lightsHandler.ToggleSwitchHandler)
	lights.Get("/:id/status", lightsHandler.GetLightSwitchStats)

	log.Fatal(app.Listen(":8080"))
}
