package utils

import (
	"gateway/internal/model"
	"log"

	"github.com/gofiber/fiber/v3"
)

func WriteErrorMessage(c fiber.Ctx,status int,  message string) error{
	return c.Status(status).JSON(model.NewErrorOutput(message))
}

func WriteErrorMessageWithLog(c fiber.Ctx,status int, message string) error{
	log.Println(message)
	return WriteErrorMessage(c, status, message)
}
