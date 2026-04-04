package server

import (
	"context"
	"encoding/json"
	"lightswitch-service/internal/core/domain"
	"lightswitch-service/internal/core/ports"
	"log"
	pr "proto/lightswitch/v1"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LightswitchServer struct{
	pr.UnimplementedLightswitchServiceServer
	service ports.LightSwitchService
	nc *nats.Conn
}

func NewLightSwitchServer(service ports.LightSwitchService, nc *nats.Conn) *LightswitchServer{
	return &LightswitchServer{service: service,nc: nc}
}

func parseUUID(id string)(*uuid.UUID, error){
	parsedId, err := uuid.Parse(id)

	if err != nil{
		return nil, status.Error(codes.InvalidArgument, "Id needs to be in uuid format")
	}

	return &parsedId, nil
}

func (s *LightswitchServer) publish(subject string, data any){
	payload, err := json.Marshal(data)
	
	if err != nil{
		log.Println("Error with publishing", subject, err)
		return
	}

	err = s.nc.Publish(subject, payload)

	if err != nil{
		log.Println("Error with publishing", subject, err)
		return
	}
}

func (s *LightswitchServer) AddLightSwitch(c context.Context,r *pr.AddLightSwitchRequest) (*pr.AddLightSwitchResponse, error) {
	lightSwitch := domain.NewLightSwitch(r.SwitchName)	

	id, err := s.service.AddLightSwitch(*lightSwitch)

	if err != nil{
		return nil, status.Error(codes.InvalidArgument, err.Error()) 
	}

	s.publish("lightswitches.added", domain.PublishAddedLightSwitch{Id: *id})

	return &pr.AddLightSwitchResponse{Id: id.String()}, nil
}

func (s *LightswitchServer) ToggleLightSwitch(c context.Context, r*pr.ToggleLightSwitchRequest) (*pr.ToggleLightSwitchResponse, error) {
	
	id, err := parseUUID(r.Id)
	
	if err != nil{
		return nil, err
	}
	
	state, err := s.service.ToggleLightSwitch(*id)

	if err != nil{
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	s.publish("lightswitches.toggled", domain.PublishToggledLightSwitch{Id: *id, State: *state})

	return &pr.ToggleLightSwitchResponse{State: *state}, nil
}

func (s *LightswitchServer) GetLightSwitchStats(c context.Context,r *pr.GetLightSwitchStatsRequest) (*pr.GetLightSwitchStatsResponse, error) {
	// TODO: implement
	panic("Not implemented")
}

func (s *LightswitchServer) GetAllLightSwitches(c context.Context, r *pr.GetAllLightSwitchesRequest) (*pr.GetAllLightSwitchesResponse, error) {
	lightSwitches, err := s.service.GetAllLightSwitches()

	if err != nil{
		return nil, status.Error(codes.Internal, "Invalid server error")
	}

	respose := pr.GetAllLightSwitchesResponse{LightSwitches: []*pr.LightSwitch{}}

	for _, lightSwitch := range *lightSwitches{
		respose.LightSwitches = append(respose.LightSwitches, &pr.LightSwitch{Id: lightSwitch.Id.String(), Name: lightSwitch.Name, State: lightSwitch.State})
	}

	return &respose, nil
}

func (s *LightswitchServer) GetLightSwitchState(c context.Context,r *pr.GetLightSwitchStateRequest) (*pr.GetLightSwitchStateResponse, error) {
	id, err := parseUUID(r.Id)

	if err != nil{
		return nil, err
	}

	state, err := s.service.GetLightSwitchState(*id)

	if err != nil{
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pr.GetLightSwitchStateResponse{State: *state}, nil
}
