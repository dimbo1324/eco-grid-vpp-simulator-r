package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/pkg/physics"
	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/utils/random"
)

// WeatherEvent represents a snapshot of environmental conditions at a specific time.
// This structure is serialized to JSON and broadcasted to Digital Twins.
type WeatherEvent struct {
	Timestamp     time.Time `json:"timestamp"`
	WindSpeed     float64   `json:"wind_speed_ms"`
	SolarRadiance float64   `json:"solar_radiance_wm2"`
	Temperature   float64   `json:"temperature_c"`
}

func main() {
	fmt.Println("Starting Environment Simulation Service (Crypto/Rand mode)...")
	fmt.Println("Press Ctrl+C to stop")

	// Simulation Configuration (mock data for MVP phase).
	const (
		SimTickRate  = 1 * time.Second // Frequency of weather updates
		WeibullShape = 2.0             // k parameter (shape)
		WeibullScale = 10.0            // lambda parameter (scale/avg speed)
		PeakSolar    = 1000.0          // Max solar irradiance at noon
		Sunrise      = 6.0             // 06:00 AM
		Sunset       = 20.0            // 08:00 PM
	)

	// Create a ticker to trigger the simulation loop at fixed intervals.
	ticker := time.NewTicker(SimTickRate)
	defer ticker.Stop()

	// Main Simulation Loop
	for t := range ticker.C {
		// 1. Generate High-Precision Random Value [0, 1)
		randVal, err := random.GenerateCanonicalFloat()
		if err != nil {
			log.Printf("Error generating random value: %v", err)
			continue
		}

		// 2. Calculate Physical Parameters
		// Calculate wind speed using Weibull distribution logic.
		windSpeed := physics.WeibullWindSpeed(randVal, WeibullShape, WeibullScale)
		// Calculate solar irradiance based on time of day.
		solarRad := physics.SolarIrradiance(t, PeakSolar, Sunrise, Sunset)

		// 3. Create Event
		event := WeatherEvent{
			Timestamp:     t,
			WindSpeed:     windSpeed,
			SolarRadiance: solarRad,
			Temperature:   20.0, // Static placeholder for MVP
		}

		// 4. Output Data (Simulating message broker publishing)
		data, _ := json.Marshal(event)
		fmt.Printf("Environment Tick: %s\n", string(data))
	}
}
