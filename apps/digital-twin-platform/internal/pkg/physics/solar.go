package physics

import (
	"math"
	"time"
)

// SolarIrradiance calculates the instantaneous solar radiation (W/m^2) based on the time of day.
// It uses a simplified sinusoidal model to simulate the sun's path between sunrise and sunset.
//
// Parameters:
//   - t: Current simulation time.
//   - peakIrradiance: Maximum possible radiation at solar noon (W/m^2).
//   - sunriseHour/sunsetHour: Time boundaries for daylight (e.g., 6.0 for 06:00).
func SolarIrradiance(t time.Time, peakIrradiance, sunriseHour, sunsetHour float64) float64 {
	// Convert current time to a fractional hour (e.g., 14:30 -> 14.5).
	currentHour := float64(t.Hour()) + float64(t.Minute())/60.0

	// Return 0 if it is night time.
	if currentHour < sunriseHour || currentHour > sunsetHour {
		return 0.0
	}

	// Calculate the sun's position using a sine wave approximation.
	dayDuration := sunsetHour - sunriseHour
	radians := (currentHour - sunriseHour) / dayDuration * math.Pi
	irradiance := peakIrradiance * math.Sin(radians)

	if irradiance < 0 {
		return 0
	}
	return irradiance
}

// SolarPanelPower calculates the electrical power output of a solar panel in Watts.
//
// Formula: P = Irradiance * Area * Efficiency
func SolarPanelPower(irradiance, area, efficiency float64) float64 {
	return irradiance * area * efficiency
}
