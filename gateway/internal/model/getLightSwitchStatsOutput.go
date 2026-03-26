package model

import "time"

type GetLightSwitchStatsOutput struct{
	TotalActiveTime time.Duration `json:"total_active_time"`
	IsActive bool `json:"is_active"`
	Name string `json:"name"`
	ActvieSince string `json:"active_since"`
}

func NewGetLightSwitchStatsOutput(lightSwitch LightSwitch) GetLightSwitchStatsOutput{


	totaActiveTime := lightSwitch.TotalFullActiveTime
	var activeSince string  

	if lightSwitch.IsActive{
		totaActiveTime += time.Since(lightSwitch.LastAcctivationTime)
		activeSince = lightSwitch.LastAcctivationTime.Format(time.RFC3339)
	}


	return GetLightSwitchStatsOutput{
		TotalActiveTime: totaActiveTime, 
		IsActive: lightSwitch.IsActive,
		Name: lightSwitch.Name,
		ActvieSince: activeSince,
	}
}
