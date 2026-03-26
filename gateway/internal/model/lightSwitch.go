package model

import "time"

type LightSwitch struct{
	Name string
	IsActive bool
	TotalFullActiveTime time.Duration // it only trakcs time after u switch off
	LastAcctivationTime time.Time
}

func NewLightSwitch(name string) *LightSwitch{
	return &LightSwitch{Name: name, IsActive: false, TotalFullActiveTime: time.Duration(0), LastAcctivationTime: time.Time{}}
}
