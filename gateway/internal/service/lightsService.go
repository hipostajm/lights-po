package service

import (
	"errors"
	"gateway/internal/model"
	"gateway/internal/repository"
	"strings"

	"github.com/google/uuid"
)

type LightsService struct{
	repository repository.LightsRepository	
}

func NewLightsService(repository repository.LightsRepository) *LightsService{
	return &LightsService{repository: repository}
}

func (s *LightsService) AddLightSwitch(swtichName string) (*uuid.UUID, error){
	
	if strings.TrimSpace(swtichName) == ""{
		return nil, errors.New("Bad name")
	}

	id, err := s.repository.AddLightSwitch(swtichName)	

	if err != nil{
		return nil, err
	}

	return id, nil
}

func (s *LightsService) ToggleLightSwitch(id uuid.UUID) (*bool,error){
	return s.repository.ToggleLightSwitch(id)
}

func (s *LightsService) GetLightSwitch(id uuid.UUID) (*model.LightSwitch, error){
	return s.repository.GetLightSwitch(id)
}

func (s *LightsService) GetLightSwitchStats(id uuid.UUID) (*model.LightSwitchStats, error){
	return s.repository.GetLightSwitchStats(id)
}

func (s *LightsService) GetAllLightSwitches() (*[]model.LightSwitch, error){
	return s.repository.GetAllLightSwitches()
}
