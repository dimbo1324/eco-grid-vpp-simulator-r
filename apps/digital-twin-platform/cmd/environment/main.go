package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/pkg/physics"
	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/utils/random"
)

type WeatherEvent struct {
	Timestamp     time.Time `json: "timestamp"`
	WindSpeed     float64   `json: "wind_speed_ms"`
	SolarRadiance float64   `json:"solar_radiance_wm2"`
	Temperature   float64   `json:"temperature_c"`
}

func main() {
	fmt.Println("Starting Environment Simulation Service (Crypto/Rand mode)...")
	fmt.Println("Press Ctrl+C to stop")
	const (
		SimTickRate  = 1 * time.Second
		WeibullShape = 2.0
		WeibullScale = 10.0
		PeakSolar    = 1000.0
		Sunrise      = 6.0
		Sunset       = 20.0
	)
	ticker := time.NewTicker(SimTickRate)
	defer ticker.Stop()
	for t := range ticker.C {
		randVal, err := random.GenerateCanonicalFloat()
		if err != nil {
			log.Printf("Error generating random value: %v", err)
			continue
		}
		windSpeed := physics.WeibullWindSpeed(randVal, WeibullShape, WeibullScale)
		solarRad := physics.SolarIrradiance(t, PeakSolar, Sunrise, Sunset)
		event := WeatherEvent{
			Timestamp:     t,
			WindSpeed:     windSpeed,
			SolarRadiance: solarRad,
			Temperature:   20.0,
		}
		data, _ := json.Marshal(event)
		fmt.Printf("Environment Tick: %s\n", string(data))
	}
}
