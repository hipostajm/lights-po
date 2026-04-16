package brodcast

import (
	"encoding/json"
	"lightswitch-service/internal/core/domain"
	"log"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type LightSwitchBrodcastImpl struct{
	mu sync.Mutex	
	subscribers map[string][]chan uuid.UUID
	client mqtt.Client
}

func NewLiLightSwitchBrodcastImpl(mu sync.Mutex, broker string, clientId string) *LightSwitchBrodcastImpl{

	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientId)

	opts.OnConnect = func(c mqtt.Client){
		log.Println("Broker connected")
	}

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal("Connection error: ", token.Error())
	}

	b := LightSwitchBrodcastImpl{mu: mu, client: client, subscribers: map[string][]chan uuid.UUID{}}

	b.client.Subscribe("lightswitches/new/confirm", 1, b.ConfirmSubscribe)

	return &b 
}

func (b *LightSwitchBrodcastImpl) Subscribe(name string) chan uuid.UUID{
	c := make(chan uuid.UUID)

	b.mu.Lock()
	defer b.mu.Unlock()

	b.subscribers[name] = append(b.subscribers[name], c)
	return  c
}

func (b *LightSwitchBrodcastImpl) Unsubscribe(name string, ch chan uuid.UUID) {
	b.mu.Lock()
	defer b.mu.Unlock()

	chans := b.subscribers[name]		

	for i,c := range chans{
		if c == ch{
			b.subscribers[name] = append(chans[:i], chans[i+1:]...)
			close(c)
			return
		}
	}
}

func  (b *LightSwitchBrodcastImpl) ConfirmSubscribe(c mqtt.Client, msg mqtt.Message){
	var data domain.ConfirmSubscribePayload

	if err := json.Unmarshal(msg.Payload(), &data); err != nil{
		log.Println(err)
		return
	}

	chans, ok := b.subscribers[data.Name]		

	if !ok{
		log.Println("Name",data.Name,"is not in process to add")
		return
	}

	id := uuid.New()

	for _, c := range chans{
		c <- id
	}
}

func (b *LightSwitchBrodcastImpl) Publish(topic string, data any) error{
	payload, err := json.Marshal(data)

	if err != nil{
		return  err
	}

	token := b.client.Publish(topic, 1, false, payload)
	
	token.Wait()

	if token.Error() != nil{
		return token.Error()
	}

	return nil
}
