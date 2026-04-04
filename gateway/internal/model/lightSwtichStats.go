package model

import "time"

type LightSwitchStats struct{
	ActiveSince time.Time `json:"active_since"`
	TotalActiveTime time.Duration `json:"total_active_time"`
}
