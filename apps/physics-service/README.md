project-root/
├── physics-service/           # Корневая папка микросервиса
│   ├── cmd/                   # Точки входа (если будет несколько способов запуска)
│   │   └── main.py            # Главный файл запуска приложения
│   │
│   ├── app/                   # Основной код приложения
│   │   ├── __init__.py
│   │   │
│   │   ├── core/              # "Чистая" бизнес-логика и математика
│   │   │   ├── simulator.py   # Класс, рассчитывающий формулы (Physics Model)
│   │   │   └── state.py       # Dataclass, описывающий текущее состояние (Pressure, Temp...)
│   │   │
│   │   ├── server/            # Сетевой слой (gRPC или API)
│   │   │   ├── grpc_server.py # Реализация gRPC сервера
│   │   │   └── protos/        # Скомпилированные .proto файлы (контракты)
│   │   │
│   │   └── settings.py        # Конфигурация (через pydantic-settings: ENV variables)
│   │
│   ├── proto/                 # Исходные .proto файлы (описание интерфейса)
│   │   └── boiler.proto       # Описание методов: GetState(), SetControls()
│   │
│   ├── tests/                 # Юнит-тесты (обязательно!)
│   ├── Dockerfile             # Инструкция для сборки образа
│   ├── requirements.txt       # Зависимости (poetry или pip)
│   └── Makefile               # Утилиты (команды типа make run, make proto)
...