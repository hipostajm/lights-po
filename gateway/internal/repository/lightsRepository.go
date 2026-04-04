package repository

import (
	"gateway/internal/model"
	"github.com/google/uuid"
)

type LightsRepository interface{
	AddLightSwitch(switchName string) (*uuid.UUID, error)
	ToggleLightSwitch(id uuid.UUID) (*bool,error)
	GetLightSwitch(id uuid.UUID) (*model.LightSwitch, error)
	GetAllLightSwitches() (*[]model.LightSwitch, error)
	GetLightSwitchStats(id uuid.UUID) (*model.LightSwitchStats, error)
}
