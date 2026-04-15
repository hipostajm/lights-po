package ports

import (
	"context"
	"lightswitch-service/internal/core/domain"

	"github.com/google/uuid"
)


type LightSwitchService interface{
	AddLightSwitch(lightSwitch domain.LightSwitch, ctx context.Context) (*uuid.UUID, error)
	ToggleLightSwitch(id uuid.UUID) (*bool, error)
	GetLightSwitch (id uuid.UUID) (*domain.LightSwitch, error)
	GetAllLightSwitches () (*[]*domain.LightSwitch, error)
	GetLightSwitchState(id uuid.UUID) (*bool, error)
	GetLightSwitchStats(id uuid.UUID) (*domain.LightSwitchStats, error)
}
