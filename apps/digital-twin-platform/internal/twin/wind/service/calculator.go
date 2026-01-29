package service

import (
	"context"
	"log"
	"time"

	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/pkg/physics"
	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/twin/domain"
	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/twin/wind/ports"
)

type WindTwinService struct {
	config    domain.TurbineConfig
	receiver  ports.WindReceiver
	publisher ports.MetricsPublisher
}

func NewWindTwinService(cfg domain.TurbineConfig, rec ports.WindReceiver, pub ports.MetricsPublisher) *WindTwinService {
	return &WindTwinService{
		config:    cfg,
		receiver:  rec,
		publisher: pub,
	}
}

func (s *WindTwinService) Start(ctx context.Context) {
	log.Printf("Digital Twin for Turbine %s started.", s.config.ID)

	area := s.config.SweptArea()

	for {
		select {
		case <-ctx.Done():
			return

		default:
			windSpeed, err := s.receiver.Receive(ctx)
			if err != nil {
				log.Printf("Error receiving wind data: %v", err)
				continue
			}

			state := s.calculateState(windSpeed, area)

			if err := s.publisher.Publish(state); err != nil {
				log.Printf("Error publishing metrics: %v", err)
			}
		}
	}
}

func (s *WindTwinService) calculateState(wind float64, area float64) domain.TurbineState {
	state := domain.TurbineState{
		Timestamp: time.Now(),
		TurbineID: s.config.ID,
		WindSpeed: wind,
	}

	switch {
	case wind < s.config.CutInSpeed:
		state.Status = "Idle (Low Wind)"
		state.PowerOutput = 0

	case wind > s.config.CutOutSpeed:
		state.Status = "Emergency Stop (High Wind)"
		state.PowerOutput = 0

	default:
		state.Status = "Active"
		power := physics.WindTurbinePower(
			physics.StandardAirDensity,
			area,
			wind,
			s.config.Efficiency,
		)
		if power > s.config.RatedPower {
			power = s.config.RatedPower
			state.Status = "Active (Rated Limit)"
		}
		state.PowerOutput = power
	}
	return state
}
