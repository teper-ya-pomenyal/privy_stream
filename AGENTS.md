# AGENTS.MD

# Project Overview
Аудио-стриминговый сервис на Go. Текущий MVP — один узел сети
(4 микросервиса: Gateway, User, Catalog, Streaming), поднимается
через единый docker-compose. Конечная цель проекта (пока не
реализована) — децентрализованная модель: независимые серверы
поднимаются разными людьми со своей музыкальной библиотекой,
а клиент хранит пул известных серверов и агрегирует по ним поиск.

## Commands
- Запуск всего проекта: docker-compose up --build
- Миграция БД: происходит автоматически при старте через сервис
  user_service_migrate (golang-migrate, путь ./user_service/migrations) для user_service
  и catalog_service_migrate (golang-migrate, путь ./catalog_service/migrations) для catalog_service
- Генерация gRPC-кода:  (protoc --go_out=. --go_opt=paths=source_relative \
                        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
                        proto/catalog/v1/catalog.proto) - для catalog_service,
                        (protoc --go_out=. --go_opt=paths=source_relative \
                        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
                        proto/user/v1/user.proto) - для user_service .

## Lint / Format
- Форматирование + управление импортами: goimports -l -w .
- Статическая проверка: go vet ./...

## Testing
Тестов пока нет в проекте. При добавлении нового кода - не полагайся на существующий тестовый прогон для валидации, а предлагай тесты как часть PR если это уместно.


## Project Structure
Монорепозиторий на go.work, 4 независимых Go-сервиса + 2 общих модуля.

- gateway/ - тонкий слой маршрутизации без бизнес-логики и БД
  - internal/handlers - HTTP-хендлеры (chi)
  - internal/clients - gRPC-клиенты к user_service/catalog_service
  - internal/middlewares - recovery/logging/CORS/auth/rate-limit
  - /keys/public.pem (том jwt_public) - публичный ключ для верификации JWT (через jwtmanager)

- user_service/ регистрация/логин, выдача JWT
  - internal/delivery, domain, repository, usecase, utils - чистая архитектура
  - /keys/private.pem (том jwt_private) - приватный ключ для подписи JWT (через jwtmanager);
    пару генерирует сервис jwt_keys_init в docker-compose при первом запуске

- catalog_service/ - каталог: треки/артисты/альбомы/плейлисты
  - internal/delivery, domain, repository, usecase, utils

- streaming_service/ - стриминг аудио через HTTP Range
  - internal/delivery, domain, storage, usecase - storage вместо repository:
    это адаптер к хранилищу файлов (сейчас ФС, позже S3/MinIO через конфиг)

Общие модули (вне отдельных сервисов, свой go.mod у каждого):
- jwtmanager/ - подпись токена (user_service) и верификация (gateway)
- proto/ - protobuf-контракты (proto/catalog/v1, proto/user/v1)

Прочее:
- api/v1/openapi.yaml - спецификация REST API Gateway
- postman/ - коллекции для ручного тестирования

## Configuration
Каждый сервис читает настройки напрямую через os.Getenv("...") в internal/config
(без сторонних библиотек вроде viper/cleanenv).
Переменные передаются через docker-compose.yml (environment:) и файл .env/.env.example в корне репозитория.

При добавлении новой настройки:
1. Добавить поле в структуру Config сервиса
2. Прочитать через os.Getenv в месте инициализации
3. Прописать переменную в .env.example
4. Прописать переменную в docker-compose.yml (environment:)


## Code Style
- context.Context — всегда первым параметром
- Обработка ошибок:
  - В domain/usecase/repository — ошибки пробрасываются как есть
    (return nil, err), без оборачивания через fmt.Errorf("%w")
  - Доменные ошибки — sentinel-значения в пакете domain
    (например, domain.ErrInvalidCredentials)
  - Маппинг на транспортный уровень (grpc) — централизованно,
    через mapDomainError() + errors.Is(), switch по sentinel-ошибкам,
    default → codes.Internal
  - Валидация формата входных данных (пустые поля и т.п.) —
    прямо в grpc-хендлере, до вызова usecase, через
    status.Error(codes.InvalidArgument, ...)
  - Валидация бизнес-правил (длина пароля и т.п.) — внутри usecase,
    через пакет internal/utils (ValidateUsername, ValidatePassword)
- Интерфейсы зависимостей — объявляются в пакете, где используются
  (не в отдельном interfaces.go и не всегда в domain)
- DI — вручную через конструкторы (New*), без DI-контейнера
- Именование файлов — по имени usecase/handler без суффиксов:
  login.go, register.go, refresh.go, logout.go (не login_usecase.go)
- Комментарии — короткие инлайн-пометки смысловых блоков внутри
  функции (// user search), не godoc-стиль над функциями/структурами
- При ошибке — всегда возвращать nil для объекта ответа, не пустую
  структуру (return nil, err), включая ошибки валидации в хендлере


## Boundaries

### Always do
- Перед завершением задачи прогонять goimports -l -w . и go vet ./...
  на затронутых модулях
- Следовать слоистой структуре delivery/domain/usecase/repository
  (или storage для streaming_service) — не смешивать бизнес-логику
  с транспортным слоем
- Новые доменные ошибки — как sentinel-значения в domain, с добавлением
  соответствующего case в mapDomainError()
- Интерфейсы объявлять в пакете, где они используются (не выносить
  в отдельный interfaces.go)
- Новые usecase — через конструктор New*, ручной DI, без контейнера
- При добавлении новой переменной окружения — обновлять и
  .env.example, и docker-compose.yml (environment:)
- Валидацию формата входных данных делать в grpc-хендлере до вызова
  usecase; бизнес-валидацию — внутри usecase

### Ask first
- Изменение схемы БД / новая миграция (в migrations/)
- Изменение .proto-контрактов в proto/ — они общие между сервисами,
  ломающие изменения затронут всех
- Добавление новой внешней зависимости в go.mod
- Изменения в jwtmanager/ — используется и в user_service, и в gateway
- Изменение структуры Config сервиса (новые обязательные поля)
- Значительный рефакторинг существующей структуры пакетов

### Never do
- Коммитить .env, keys/private.pem, keys/public.pem или любые
  секреты/ключи
- Менять уже применённые (не последние) .up.sql/.down.sql файлы
  миграций задним числом — только новые миграции
- Добавлять golangci-lint или менять CI без явной просьбы (пока не
  дошли до этого этапа по плану)
- Писать тесты «для галочки» без реальной проверки логики только
  ради наличия покрытия
