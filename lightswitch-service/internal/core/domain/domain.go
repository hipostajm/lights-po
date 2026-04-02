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
