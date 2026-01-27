package domain

import "time"

// WeatherEvent represents a snapshot of environmental conditions at a specific time.
// This structure is serialized to JSON and broadcasted to Digital Twins.
type WeatherEvent struct {
	Timestamp     time.Time `json:"timestamp"`
	WindSpeed     float64   `json:"wind_speed_ms"`
	SolarRadiance float64   `json:"solar_radiance_wm2"`
	Temperature   float64   `json:"temperature_c"`
}
