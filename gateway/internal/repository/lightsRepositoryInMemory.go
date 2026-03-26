package repository

import (
	"errors"
	"gateway/internal/model"
	"time"

	"github.com/google/uuid"
)

type LightsRepositoryInMemory struct{
	switches map[uuid.UUID]*model.LightSwitch
}

func NewLightsRepositoryInMemory() *LightsRepositoryInMemory{
	return &LightsRepositoryInMemory{switches: make(map[uuid.UUID]*model.LightSwitch)}	
}


func (r *LightsRepositoryInMemory) AddLightSwitch(switchName string) (*uuid.UUID, error){
	var id uuid.UUID
	for {
		id = uuid.New()
		if _, ok := r.switches[id]; !ok{
			break	
		}
	}
	r.switches[id] = model.NewLightSwitch(switchName)
	return &id, nil
}

func (r *LightsRepositoryInMemory) ToggleLightSwitch(id uuid.UUID) (*bool,error){

	lightSwitch, err := r.GetLightSwitch(id)

	if err != nil{
		return nil,	err
	}


	if lightSwitch.IsActive{
		lightSwitch.TotalFullActiveTime += time.Since(lightSwitch.LastAcctivationTime)
	} else{
		lightSwitch.LastAcctivationTime = time.Now()
	}

	lightSwitch.IsActive = !lightSwitch.IsActive

	return &lightSwitch.IsActive, nil
}

func (r *LightsRepositoryInMemory) GetLightSwitch(id uuid.UUID) (*model.LightSwitch, error){
	
	lightSwitch, ok := r.switches[id]

	if !ok{
		return nil,errors.New("Id not found")
	}

	return lightSwitch, nil
}
