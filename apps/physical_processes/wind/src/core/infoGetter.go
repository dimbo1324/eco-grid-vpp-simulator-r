package core

import (
	"strings"
	"time"
)

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

// func generateWeatherSimulation(loc interfaces.LocationProfile, t time.Time) interfaces.EnvironmentState {
// 	season := getSeason(t.Month())
// 	timeOfDay := getTimeOfDay(t.Hour())

// 	var baseTemp float64
// 	switch season {
// 	case properties["win"]:
// 		baseTemp = -2.0
// 	case properties["spr"]:
// 		baseTemp = 12.0
// 	case properties["sum"]:
// 		baseTemp = 25.0
// 	case properties["aut"]:
// 		baseTemp = 14.0
// 	}

// 	if timeOfDay == properties["mor"] || timeOfDay == properties["nig"] {
// 		baseTemp -= 5.0
// 	}

// 	scenarios := []interfaces.WeatherScenario{
// 		{"Ясно", 2, 500, 5, 1},             // Теплее, высокое давление, средний ветер
// 		{"Облачно", -1, 0, 6, 2},           // Чуть прохладнее, стандартное давление
// 		{"Дождь", -3, -500, 8, 3},          // Прохладно, низкое давление, ветер сильнее
// 		{"Сильный Ветер", -2, -800, 12, 3}, // Ветрено
// 		{"Шторм", -5, -1500, 15, 5},        // Холодно, очень низкое давление, сильный ветер, турбулентность
// 		{"Штиль", 1, 200, 2, 1},            // Спокойно
// 	}

// 	// i, _ := randomers.randomInt()
// }
