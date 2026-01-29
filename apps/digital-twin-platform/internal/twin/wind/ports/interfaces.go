package ports

import (
	"context"

	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/twin/domain"
)

type WindReceiver interface {
	Receive(ctx context.Context) (float64, error)
}

type MetricsPublisher interface {
	Publish(state domain.TurbineState) error
}
