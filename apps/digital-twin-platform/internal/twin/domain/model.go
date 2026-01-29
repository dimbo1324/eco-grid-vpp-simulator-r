package domain

import (
	"time"

	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/constants"
)

type TurbineConfig struct {
	ID            string
	RotorDiameter float64
	RatedPower    float64
	CutInSpeed    float64
	CutOutSpeed   float64
	Efficiency    float64
}

type TurbineState struct {
	Timestamp   time.Time `json:"timestamp"`
	TurbineID   string    `json:"turbine_id"`
	WindSpeed   float64   `json:"wind_speed_ms"`
	PowerOutput float64   `json:"power_output_w"`
	Status      string    `json:"status"`
}

func (c TurbineConfig) SweptArea() float64 {
	radius := 0.5 * c.RotorDiameter
	return constants.Pi * radius * radius
}
