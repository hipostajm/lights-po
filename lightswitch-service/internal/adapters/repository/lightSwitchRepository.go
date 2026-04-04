package repository

import (
	"errors"
	"lightswitch-service/internal/core/domain"
	"maps"
	"slices"

	"github.com/google/uuid"
)

type LightSwitchInMemoryRepository struct{
	lighSwitches map[uuid.UUID]*domain.LightSwitch
}

func NewLightSwitchInMemoryRepository() *LightSwitchInMemoryRepository {
	return &LightSwitchInMemoryRepository{lighSwitches: make(map[uuid.UUID]*domain.LightSwitch)}
}

func (r *LightSwitchInMemoryRepository) AddLightSwitch(lightSwitch domain.LightSwitch) (*uuid.UUID, error){
	var id uuid.UUID
	
	for{
		id = uuid.New()
		if _, ok := r.lighSwitches[id]; !ok{
			break
		}
	}

	lightSwitch.Id = id
	r.lighSwitches[id] = &lightSwitch

	return &id, nil
}

func (r *LightSwitchInMemoryRepository) ToggleLightSwitch(id uuid.UUID) (*bool, error){
	
	lightSwitch, ok := r.lighSwitches[id]
	
	if !ok{
		return nil, errors.New("Id not found")
	}


	lightSwitch.State = !lightSwitch.State

	return &lightSwitch.State, nil
}

func (r *LightSwitchInMemoryRepository) GetLightSwitch(id uuid.UUID) (*domain.LightSwitch, error){

	lightSwitch, ok := r.lighSwitches[id]

	if !ok{
		return nil, errors.New("Id not found")
	}

	return lightSwitch, nil
}

func (r *LightSwitchInMemoryRepository) GetAllLightSwitches() (*[]*domain.LightSwitch, error){
	lightSwitches := slices.Collect(maps.Values(r.lighSwitches))
	return &lightSwitches, nil
}
