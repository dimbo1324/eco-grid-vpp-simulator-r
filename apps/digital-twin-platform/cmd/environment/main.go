package main

import (
	"fmt"

	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/environment/adapters"
	"github.com/dimbo1324/eco-grid-vpp-simulator-r/internal/environment/service"
)

func main() {
	fmt.Println("Starting Environment Simulation Service (Crypto/Rand mode)...")
	fmt.Println("Press Ctrl+C to stop")

	producer := adapters.NewConsoleProducer()
	simService := service.NewSimulator(producer)
	simService.Start()
}
