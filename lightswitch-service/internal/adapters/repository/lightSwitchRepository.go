package repository

import (
	"errors"
	"lightswitch-service/internal/core/domain"
	"maps"
	"slices"
	"time"

	"github.com/google/uuid"
)

type LightSwitchInMemoryRepository struct{
	lighSwitches map[uuid.UUID]*domain.LightSwitch
}

func NewLightSwitchInMemoryRepository() *LightSwitchInMemoryRepository {
	return &LightSwitchInMemoryRepository{lighSwitches: make(map[uuid.UUID]*domain.LightSwitch)}
}

func (r *LightSwitchInMemoryRepository) AddLightSwitch(lightSwitch domain.LightSwitch) (error){

	_, ok := r.lighSwitches[lightSwitch.Id]

	if ok{
		return errors.New("Id is used")
	}

	r.lighSwitches[lightSwitch.Id] = &lightSwitch
	return nil
}

func (r *LightSwitchInMemoryRepository) ToggleLightSwitch(id uuid.UUID) (*bool, error){
	
	lightSwitch, ok := r.lighSwitches[id]
	
	if !ok{
		return nil, errors.New("Id not found")
	}


	lightSwitch.State = !lightSwitch.State

	if lightSwitch.State{
		lightSwitch.LastAcctivationTime = time.Now()
	}

	return &lightSwitch.State, nil
}

func (r *LightSwitchInMemoryRepository) GetLightSwitch(id uuid.UUID) (*domain.LightSwitch, error){

	lightSwitch, ok := r.lighSwitches[id]

	if !ok{
		return nil, errors.New("Id not found")
	}

	return lightSwitch, nil
}

func (r *LightSwitchInMemoryRepository) GetLightSwitchByName(name string) (*domain.LightSwitch, error){
	for _, ls := range r.lighSwitches{
		if ls.Name == name{
			return ls, nil
		}	
	}
	return nil, errors.New("Name not found")
}

func (r *LightSwitchInMemoryRepository) GetAllLightSwitches() (*[]*domain.LightSwitch, error){
	lightSwitches := slices.Collect(maps.Values(r.lighSwitches))
	return &lightSwitches, nil
}
