package model

import "github.com/google/uuid"

type AddLightSwitchOutput struct{
	Id uuid.UUID `json:"id"`
}

func NewAddLightSwitchOutput(id uuid.UUID) AddLightSwitchOutput{
	return AddLightSwitchOutput{Id: id}
}
