package ports

import (
	"lightswitch-service/internal/core/domain"

	"github.com/google/uuid"
)

type LightSwitchRepository interface{
	AddLightSwitch(lightSwitch domain.LightSwitch) (error)
	ToggleLightSwitch(id uuid.UUID) (*bool, error)
	GetLightSwitch (id uuid.UUID) (*domain.LightSwitch, error)
	GetAllLightSwitches () (*[]*domain.LightSwitch, error)
	GetLightSwitchByName(name string) (*domain.LightSwitch, error)
}

