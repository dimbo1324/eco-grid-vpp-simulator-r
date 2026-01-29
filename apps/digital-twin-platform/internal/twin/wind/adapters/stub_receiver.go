package adapters

import (
	"context"
	"time"

	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/utils/random"
)

type StubWindReceiver struct {
	// TODO: Можно добавить настройки задержки
}

func NewStubWindReceiver() *StubWindReceiver {
	return &StubWindReceiver{}
}

func (s *StubWindReceiver) Receive(ctx context.Context) (float64, error) {
	select {
	case <-time.After(1 * time.Second):
		b, e := 0.0, 30.0
		wind, err := random.CreateRandFloat(b, e)
		if err != nil {
			return 0.0, err
		}
		return wind, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}
