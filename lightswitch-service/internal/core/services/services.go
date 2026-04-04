package services

import (
	"lightswitch-service/internal/core/domain"
	"lightswitch-service/internal/core/ports"
	"time"

	"github.com/google/uuid"
)

type LightSwitchService struct {
	repostitory ports.LightSwitchRepository
}

func NewLightSwithcService(repo ports.LightSwitchRepository) *LightSwitchService{
	return &LightSwitchService{repostitory: repo}
}

func (s *LightSwitchService)AddLightSwitch(lightSwitch domain.LightSwitch) (*uuid.UUID, error){
	return s.repostitory.AddLightSwitch(lightSwitch)	
}

func (s *LightSwitchService)ToggleLightSwitch(id uuid.UUID) (*bool, error){
	return s.repostitory.ToggleLightSwitch(id)
}

func (s *LightSwitchService)GetLightSwitch (id uuid.UUID) (*domain.LightSwitch, error){
	return s.repostitory.GetLightSwitch(id)
}

func (s *LightSwitchService)GetAllLightSwitches() (*[]*domain.LightSwitch, error){
	return s.repostitory.GetAllLightSwitches()
}

func (s *LightSwitchService)GetLightSwitchState(id uuid.UUID) (*bool, error){
	return s.GetLightSwitchState(id)	
}


func (s *LightSwitchService)GetLightSwitchStats(id uuid.UUID) (*domain.LightSwitchStats, error){
	lightSwitch, err := s.repostitory.GetLightSwitch(id)

	if err != nil{
		return nil, err
	}

	var totalActiveTime time.Duration
	
	if lightSwitch.LastAcctivationTime.IsZero(){
		totalActiveTime = 0;
	} else{
		totalActiveTime = lightSwitch.TotalClosedActiveTime+time.Since(lightSwitch.LastAcctivationTime)
	}


	return &domain.LightSwitchStats{TotalActiveTime: totalActiveTime,  ActiveSine: lightSwitch.LastAcctivationTime}, nil
}
