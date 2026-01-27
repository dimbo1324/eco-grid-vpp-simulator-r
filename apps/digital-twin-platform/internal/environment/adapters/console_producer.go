package adapters

import (
	"encoding/json"
	"fmt"

	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/environment/domain"
)

type ConsoleProducer struct{}

func NewConsoleProducer() *ConsoleProducer {
	return &ConsoleProducer{}
}

func (p *ConsoleProducer) Publish(event domain.WeatherEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	fmt.Printf("Environment Tick: %s\n", string(data))
	return nil
}
