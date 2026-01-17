package interfaces

import "time"

// TODO: docs and discription
type LocationProfile struct {
	ID         int
	City       string
	Region     string
	Country    string
	Latitude   float64
	Longitude  float64
	Timezone   string
	BaseHeight float64
	HubHeight  float64
	Roughness  float64
}

// TODO: docs and discription
type EnvironmentState struct {
	Season      string
	TimeOfDay   string
	WeatherType string
	Temperature float64
	Pressure    float64
	WeibullK    float64
	WeibullC    float64
	Timestamp   time.Time
}

// TODO: docs and discription
type WeatherScenario struct {
	Name        string
	TempOffset  int64
	PressOffset int64
	WindScale   int64
	Turbulence  int64
}
