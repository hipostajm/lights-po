package domain

import (
	"time"

	"github.com/google/uuid"
)

type LightSwitch struct{
	Id uuid.UUID
	Name string
	State bool
	TotalClosedActiveTime time.Duration 
	LastAcctivationTime time.Time
}

func NewLightSwitch(name string) *LightSwitch{
	return &LightSwitch{Name: name, State: false, TotalClosedActiveTime: 0, LastAcctivationTime: time.Time{}}
}

type PublishAddedLightSwitch struct{
	Id uuid.UUID `json:"id"`
}

type PublishToggledLightSwitch struct{
	Id uuid.UUID `json:"id"`
	State bool `json"state"`
}

type LightSwitchStats struct{
	TotalActiveTime time.Duration
	ActiveSine time.Time
}

type ConfirmSubscribePayload struct{
	Name string `json:"name"`
}

type ToggleLightSwitchPayload struct{
	Name string `json:"name"`
	State bool `json:"state"`
}

type NewLightSwitchPayload struct{
	Name string `json:"name"`
}
