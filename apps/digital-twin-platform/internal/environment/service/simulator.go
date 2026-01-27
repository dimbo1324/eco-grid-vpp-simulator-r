package service

import (
	"log"
	"time"

	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/environment/domain"
	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/pkg/physics"
	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/utils/random"
)

type EventProducer interface {
	Publish(event domain.WeatherEvent) error
}

type Simulator struct {
	producer EventProducer
}

func NewSimulator(p EventProducer) *Simulator {
	return &Simulator{producer: p}
}

func (s *Simulator) Start() {

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
		event := domain.WeatherEvent{
			Timestamp:     t,
			WindSpeed:     windSpeed,
			SolarRadiance: solarRad,
			Temperature:   20.0, // Static placeholder for MVP
		}

		// 4. Output Data (Simulating message broker publishing)
		if err := s.producer.Publish(event); err != nil {
			log.Printf("Failed to publish event: %v", err)
		}
	}
}
