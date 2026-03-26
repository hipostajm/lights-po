package handler

import (
	"errors"
	"gateway/internal/model"
	"gateway/internal/service"
	"gateway/internal/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type LightsHandler struct {
	service *service.LightsService
}

func NewLightsHandler(service *service.LightsService) LightsHandler{
	return LightsHandler{service: service}
}

func getIdFromPathParam(c fiber.Ctx) (*uuid.UUID, error){
	stringId := c.Params("id")

	if stringId == ""{
		return nil, errors.New("Id can not be empty")
	}

	id, err := uuid.Parse(stringId)

	if err != nil{
		return nil, errors.New("Bad uuid format")
	}
	return &id, nil
}

func (h *LightsHandler) AddLightSwitch(c fiber.Ctx) error{
	
	var input model.AddLightSwitchInput 

	if err := c.Bind().Body(&input); err != nil{
		return utils.WriteErrorMessageWithLog(c,fiber.StatusBadRequest,"Bad json")
	}

	id, err := h.service.AddLightSwitch(input.Name)

	if err != nil{
		return utils.WriteErrorMessageWithLog(c, fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(model.NewAddLightSwitchOutput(*id))
}


func (h *LightsHandler) ToggleSwitchHandler(c fiber.Ctx) error{
	

	id, err := getIdFromPathParam(c)

	if err != nil{
		return utils.WriteErrorMessageWithLog(c, fiber.StatusBadRequest, err.Error())
	}

	state, err := h.service.ToggleLightSwitch(*id)

	if err != nil{
		return utils.WriteErrorMessageWithLog(c, fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(model.NewToggleLightSwitchOutput(*state))
}

func (h *LightsHandler) GetLightSwitchStats(c fiber.Ctx) error{
	id, err := getIdFromPathParam(c)

	if err != nil{
		return utils.WriteErrorMessageWithLog(c, fiber.StatusBadRequest, err.Error())
	}	

	lightSwitch, err := h.service.GetLightSwitch(*id)	

	if err != nil{
		return utils.WriteErrorMessageWithLog(c, fiber.StatusBadRequest, err.Error())
	}
	

	return c.Status(fiber.StatusOK).JSON(model.NewGetLightSwitchStatsOutput(*lightSwitch))
}

