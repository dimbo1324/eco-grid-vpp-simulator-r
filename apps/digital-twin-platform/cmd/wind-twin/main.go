package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/twin/domain"
	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/twin/wind/adapters"
	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/twin/wind/service"
)

func main() {
	cfg := domain.TurbineConfig{
		ID:            "WT-01-Alpha",
		RotorDiameter: 164.0,
		RatedPower:    8000000,
		CutInSpeed:    3.0,
		CutOutSpeed:   25.0,
		Efficiency:    0.45,
	}
	receiver := adapters.NewStubWindReceiver()
	publisher := adapters.NewStdoutPublisher()
	twinService := service.NewWindTwinService(cfg, receiver, publisher)
	ctx, cancel := context.WithCancel(context.Background())

	// TODO: Необходимо рассмотреть целесообразность приведенной ниже логики и подумать над оптимальным использованием ресурсов машины, возможно, стоит реализовать решение на более низком уровне.
	go twinService.Start(ctx)
	fmt.Println("Wind Turbine Digital Twin is running. Press Ctrl+C to stop.")
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan
	fmt.Println("\nShutting down...")
	cancel()
}
