package model

import (
	"github.com/google/uuid"
)

type LightSwitch struct{
	Id uuid.UUID
	Name string
	State bool
}

func NewLightSwitch(name string) *LightSwitch{
	return &LightSwitch{Name: name, State: false}
}
