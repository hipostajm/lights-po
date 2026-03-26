package model

type ToggleLightSwitchOutput struct{
	State bool `json:"state"`
}

func NewToggleLightSwitchOutput(state bool)ToggleLightSwitchOutput{
	return ToggleLightSwitchOutput{State: state}
}
