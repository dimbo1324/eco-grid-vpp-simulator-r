package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"wind/src/core"
	"wind/src/core/interfaces"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Wind Digital Twin Simulation Engine [v1.0] ===")
	fmt.Println("Выберите локацию для симуляции:")
	fmt.Println("1. New York, USA")
	fmt.Print("Ввод > ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "1" {
		loc := interfaces.Locations[1]
		runSimulation(loc)
	} else {
		fmt.Println("Ошибка: Локация не найдена. Запускаем New York по умолчанию.")
		runSimulation(interfaces.Locations[1])
	}
}

func runSimulation(loc interfaces.LocationProfile) {
	location, err := time.LoadLocation(loc.Timezone)
	if err != nil {
		fmt.Printf("Внимание: таймзона %s не найдена, используем UTC.\n", loc.Timezone)
		location = time.UTC
	}

	localTime := time.Now().In(location)

	env := core.GenerateWeatherSimulation(loc, localTime)

	fmt.Println("\n---------------------------------------------------")
	fmt.Printf("ОТЧЕТ СИМУЛЯЦИИ: %s, %s\n", loc.City, loc.Country)
	fmt.Println("---------------------------------------------------")

	fmt.Printf("Координаты:      %.4f, %.4f\n", loc.Latitude, loc.Longitude)
	fmt.Printf("Дата:            %d %s (%s)\n",
		env.Timestamp.Day(), env.Timestamp.Month().String(), env.Season)
	fmt.Printf("clock  Время:           %s (%s)\n",
		env.Timestamp.Format("15:04:05"), env.TimeOfDay)

	fmt.Println("\n--- Погодные условия (Динамика) ---")
	fmt.Printf("Состояние:          %s\n", env.WeatherType)
	fmt.Printf("Температура:        %.2f °C\n", env.Temperature)
	fmt.Printf("Атм. Давление:      %.0f Pa\n", env.Pressure)

	fmt.Println("\n--- Параметры Ветра (Физика модели) ---")
	params := map[string]interface{}{
		"Weibull Shape (k)": fmt.Sprintf("%.2f (Характер порывов)", env.WeibullK),
		"Weibull Scale (c)": fmt.Sprintf("%.2f м/с (Ср. скорость)", env.WeibullC),
		"Roughness (alpha)": loc.Roughness,
		"Hub Height":        fmt.Sprintf("%.1f м", loc.HubHeight),
		"Base Height":       fmt.Sprintf("%.1f м", loc.BaseHeight),
	}

	for key, val := range params {
		fmt.Printf("• %-20s : %v\n", key, val)
	}
	fmt.Println("---------------------------------------------------")
	fmt.Println("Система готова к расчету выработки энергии...")
}
