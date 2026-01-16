package interfaces

import "time"

// Описывает статические данные локации (то, что не меняется).
type LocationProfile struct {
	ID         int
	City       string
	Region     string
	Country    string
	Latitude   float64
	Longitude  float64
	Timezone   string
	BaseHeight float64 // Высота измерения (обычно 10м)
	HubHeight  float64 // Высота ступицы турбины (например, 100м)
	Roughness  float64 // Шероховатость (0.1 море, 0.4 город)
}

// Описывает текущий "срез" реальности.
type EnvironmentState struct {
	Season      string    // Зима, Весна...
	TimeOfDay   string    // Утро, День, Вечер, Ночь
	WeatherType string    // Ясно, Дождь, Шторм...
	Temperature float64   // Градусы Цельсия
	Pressure    float64   // Паскали (Pa)
	WeibullK    float64   // Параметр формы (Shape)
	WeibullC    float64   // Параметр масштаба (Scale)
	Timestamp   time.Time // Локальное время
}

// Помогает задать границы параметров для разных типов погоды.
type WeatherScenario struct {
	Name        string
	TempOffset  int64 // Смещение температуры относительно сезона
	PressOffset int64 // Смещение давления (шторм = низкое, ясно = высокое)
	WindScale   int64 // Базовая сила ветра (для Weibull C)
	Turbulence  int64 // Уровень турбулентности (влияет на Weibull K: выше число = ниже K)
}
