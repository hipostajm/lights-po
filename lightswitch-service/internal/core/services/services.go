package services

import (
	"context"
	"errors"
	"lightswitch-service/internal/core/domain"
	"lightswitch-service/internal/core/ports"
	"time"
	"github.com/google/uuid"
)

type LightSwitchService struct {
	repostitory ports.LightSwitchRepository
	brodact ports.LightSwitchBrodcast
	waitTime time.Duration
}

func NewLightSwithcService(repo ports.LightSwitchRepository, brodact ports.LightSwitchBrodcast, waitTime time.Duration) *LightSwitchService{
	return &LightSwitchService{repostitory: repo, brodact: brodact, waitTime: waitTime}
}

func (s *LightSwitchService)AddLightSwitch(lightSwitch domain.LightSwitch) (*uuid.UUID, error){

	_, err :=  s.repostitory.GetLightSwitchByName(lightSwitch.Name)

	if err == nil{
		return nil, errors.New("Name is taken")
	}

	ch := s.brodact.Subscribe(lightSwitch.Name)
	defer s.brodact.Unsubscribe(lightSwitch.Name,ch)

	ctx, cancel := context.WithTimeout(context.Background(), s.waitTime)
	defer cancel()

	err = s.brodact.Publish("lightswitches/new", domain.NewLightSwitchPayload{Name: lightSwitch.Name})

	if err != nil{
		return nil, err
	}

	select{
	case id := <- ch:
		lightSwitch.Id = id
		return &id, s.repostitory.AddLightSwitch(lightSwitch)	
	case <- ctx.Done():
		return nil, errors.New("Time window passed")
	}
}

func (s *LightSwitchService)ToggleLightSwitch(id uuid.UUID) (*bool, error){

	state, err := s.repostitory.ToggleLightSwitch(id)

	if err != nil{
		return nil, err
	}

	ls, err := s.repostitory.GetLightSwitch(id)


	if err != nil{
		return nil, err
	}

	err = s.brodact.Publish("lightswitches/toggle", domain.ToggleLightSwitchPayload{State: *state, Name: ls.Name})
	
	if err != nil{
		return nil, err
	}

	return state, nil
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
