package ports

import "github.com/google/uuid"

type LightSwitchBrodcast interface{
	Publish(topic string, data any) error
	Subscribe(name string) chan uuid.UUID
	Unsubscribe(name string, ch chan uuid.UUID)
}
