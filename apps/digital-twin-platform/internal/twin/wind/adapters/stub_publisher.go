package adapters

import (
	"encoding/json"
	"fmt"

	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/twin/domain"
)

type StdoutPublisher struct{}

func NewStdoutPublisher() *StdoutPublisher {
	return &StdoutPublisher{}
}

func (p *StdoutPublisher) Publish(state domain.TurbineState) error {
	data, _ := json.Marshal(state)
	fmt.Printf("[WindTwin-%s] Calculated: %s\n", state.TurbineID, string(data))
	return nil
}
