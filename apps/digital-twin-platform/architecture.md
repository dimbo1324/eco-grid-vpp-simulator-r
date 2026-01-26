/digital-twin-platform
├── /cmd                    # Точки входа (Main файлы)
│   ├── /environment        # main.go для симулятора среды
│   ├── /wind-twin          # main.go для турбины
│   └── /solar-twin         # main.go для панели
│
├── /internal               # Приватный код (Бизнес-логика)
│   ├── /pkg                # Общие библиотеки (Physics Core)
│   │   ├── /physics        # Формулы (Вейбул, инсоляция) - чтобы не дублировать!
│   │   └── /transport      # Обертки над NATS/gRPC
│   │
│   ├── /environment        # Логика модуля Environment
│   │   ├── /domain         # Сущности (WeatherState, Location)
│   │   ├── /adapters       # Источники (RandomProvider, OpenWeatherProvider)
│   │   └── /service        # Основной цикл генерации
│   │
│   └── /twin               # Логика Двойников (можно переиспользовать core)
│       ├── /domain         # Сущности (TurbineState, PowerCurve)
│       └── /service        # Логика расчета выработки
│
├── /configs                # Конфигурации (YAML)
├── docker-compose.yml      # Поднимает NATS, InfluxDB и сервисы
└── go.mod                  # Единый модуль (или workspace)