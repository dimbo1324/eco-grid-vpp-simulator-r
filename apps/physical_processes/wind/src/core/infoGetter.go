package core

import (
	"strings"
	"time"
	"wind/internal/utils/randomers"
	"wind/src/core/interfaces"
)

const mathPrec int = 100

var properties = map[string]string{
	"win":    "Winter",
	"spr":    "Spring",
	"sum":    "Summer",
	"aut":    "Autumn",
	"mor":    "Morning",
	"aft":    "Afternoon",
	"eve":    "Evening",
	"nig":    "Night",
	"defVal": strings.ToUpper("unknown"),
}

var scenarios = [...]interfaces.WeatherScenario{
	{Name: "Ясно", TempOffset: 2, PressOffset: 500, WindScale: 5, Turbulence: 1},             // Теплее, высокое давление, средний ветер
	{Name: "Облачно", TempOffset: -1, PressOffset: 0, WindScale: 6, Turbulence: 2},           // Чуть прохладнее, стандартное давление
	{Name: "Дождь", TempOffset: -3, PressOffset: -500, WindScale: 8, Turbulence: 3},          // Прохладно, низкое давление, ветер сильнее
	{Name: "Сильный Ветер", TempOffset: -2, PressOffset: -800, WindScale: 12, Turbulence: 3}, // Ветрено
	{Name: "Шторм", TempOffset: -5, PressOffset: -1500, WindScale: 15, Turbulence: 5},        // Холодно, очень низкое давление, сильный ветер, турбулентность
	{Name: "Штиль", TempOffset: 1, PressOffset: 200, WindScale: 2, Turbulence: 1},            // Спокойно
}

// TODO: docs and discription
func getSeason(month time.Month) string {
	switch month {
	case time.December, time.January, time.February:
		return properties["win"]
	case time.March, time.April, time.May:
		return properties["spr"]
	case time.June, time.July, time.August:
		return properties["sum"]
	case time.September, time.October, time.November:
		return properties["aut"]
	default:
		return properties["defVal"]
	}
}

// TODO: docs and discription
func getTimeOfDay(hour int) string {
	switch {
	case hour >= 5 && hour < 12:
		return properties["mor"]
	case hour >= 12 && hour < 17:
		return properties["aft"]
	case hour >= 17 && hour < 22:
		return properties["eve"]
	default:
		return properties["nig"]
	}
}

// TODO: docs and discription
func generateWeatherSimulation( /*loc interfaces.LocationProfile,*/ t time.Time) interfaces.EnvironmentState {
	season := getSeason(t.Month())
	timeOfDay := getTimeOfDay(t.Hour())

	var baseTemp float64
	switch season {
	case properties["win"]:
		baseTemp = -2.0
	case properties["spr"]:
		baseTemp = 12.0
	case properties["sum"]:
		baseTemp = 25.0
	case properties["aut"]:
		baseTemp = 14.0
	}

	if timeOfDay == properties["mor"] || timeOfDay == properties["nig"] {
		baseTemp -= 5.0
	}

	end := int64(len(scenarios) - 1)

	i := randomers.GetRandInt(0, end)
	currentScenario := scenarios[i]

	noiseT := randomers.GetRandFloat(-2.0, 2.0, mathPrec)
	finalTemp := baseTemp + float64(currentScenario.TempOffset) + noiseT

	noiseP := randomers.GetRandFloat(-100.0, 100.0, mathPrec)
	finalPress := 101325.0 + float64(currentScenario.PressOffset) + noiseP

	noiseC := randomers.GetRandFloat(-1.0, 1.0, mathPrec)
	weibullC := float64(currentScenario.WindScale) + noiseC
	if weibullC < 1 {
		weibullC = 1
	}

	baseK := 2.2
	turbulenceFactor := float64(currentScenario.Turbulence) * 0.15
	noiseK := randomers.GetRandFloat(-0.1, 0.1, mathPrec)
	weibullK := baseK - turbulenceFactor + noiseK

	return interfaces.EnvironmentState{
		Season:      season,
		TimeOfDay:   timeOfDay,
		WeatherType: currentScenario.Name,
		Temperature: finalTemp,
		Pressure:    finalPress,
		WeibullK:    weibullK,
		WeibullC:    weibullC,
		Timestamp:   t,
	}
}
