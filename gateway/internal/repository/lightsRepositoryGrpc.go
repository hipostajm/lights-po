package repository

import (
	"context"
	"gateway/internal/model"
	lightswitchv1 "proto/lightswitch/v1"
	"time"

	"github.com/google/uuid"
)

type LightsRepositoryGrpc struct{
	client lightswitchv1.LightswitchServiceClient
	ctxTime time.Duration
}

func NewLightsRepositoryGrpc(client lightswitchv1.LightswitchServiceClient,ctxTime time.Duration) *LightsRepositoryGrpc{
	return &LightsRepositoryGrpc{client: client, ctxTime: ctxTime}
}

func (r *LightsRepositoryGrpc)AddLightSwitch(switchName string) (*uuid.UUID, error){
	ctx, cancel := r.getCtxAndCancel()
	defer cancel()

	response, err := r.client.AddLightSwitch(ctx, &lightswitchv1.AddLightSwitchRequest{SwitchName: switchName})
	if err != nil{
		return nil, err
	}
	id, err := uuid.Parse(response.Id)
	if err != nil{
		return nil, err
	}
	return &id, nil
}

func (r *LightsRepositoryGrpc)ToggleLightSwitch(id uuid.UUID) (*bool,error){
	ctx, cancel := r.getCtxAndCancel()
	defer cancel()
	
	response, err := r.client.ToggleLightSwitch(ctx, &lightswitchv1.ToggleLightSwitchRequest{Id: id.String()})

	if err != nil{
		return nil,err 
	}
	
	return &response.State, nil
}


func (r *LightsRepositoryGrpc)GetLightSwitch(id uuid.UUID) (*model.LightSwitch, error){
	ctx, cancel := r.getCtxAndCancel()
	defer cancel()

	response, err := r.client.GetLightSwitch(ctx, &lightswitchv1.GetLightSwitchRequest{Id: id.String()})
	
	if err != nil{
		return nil, err
	}

	lightSwitch := model.NewLightSwitch(response.LightSwitch.Name)
	lightSwitch.State = response.LightSwitch.State
	lightSwitch.Id = id

	return lightSwitch, nil
}


func (r *LightsRepositoryGrpc) GetAllLightSwitches() (*[]model.LightSwitch, error){
	ctx, cancel := r.getCtxAndCancel()
	defer cancel()

	response, err := r.client.GetAllLightSwitches(ctx, &lightswitchv1.GetAllLightSwitchesRequest{})

	if err != nil{
		return nil, err
	}

	lightSwitches := []model.LightSwitch{}
	
	for _, responseLS := range response.LightSwitches{

		ls, err := mapLightSwitchFromRespons(*responseLS)

		if err != nil{
			return nil, err
		}

		lightSwitches = append(lightSwitches, *ls)
	}

	return &lightSwitches, nil
}

func (r *LightsRepositoryGrpc) getCtxAndCancel() (context.Context, context.CancelFunc){
	return context.WithTimeout(context.Background(), r.ctxTime)
}

func mapLightSwitchFromRespons(reqSL lightswitchv1.LightSwitch) (*model.LightSwitch, error){
	ls := model.LightSwitch{}

	ls.Name = reqSL.Name
	ls.State = reqSL.State

	id, err := uuid.Parse(reqSL.Id)

	if err != nil{
		return nil ,err
	}

	ls.Id = id

	return &ls, nil
}
