# Orders Management Service
## gRPC сервис управления пунктом выдачи заказов на языке Go.
### В имплементации также используются PostgreSQL, Apache Kafka, Redis, Prometheus и Grafana.

### Список команд

1. **Принять заказ от курьера** [Orders.AcceptOrder]
2. **Вернуть заказ курьеру** [Orders.ReturnOrder]
3. **Выдать заказы клиенту** [Orders.DeliverOrders]
4. **Получить список заказов** [Orders.ListOrders]
5. **Принять возврат от клиента** [Returns.AcceptReturn]
6. **Получить список возвратов** [Returns.ListReturns]

### Руководство по работе с проектом

#### Настройка окружения

1. Создайте файл `.env` в корне проекта и добавьте необходимые переменные окружения. В качестве файла `.env` можно использовать оригинальный
2. Установите бинарные зависимости через `make deps`
3. Соберите Docker-контейнеры через `make build`.
4. Накатайте миграции через `make migrate`.
5. Запустите сервер с через `make server`. Пример возможных сообщений можно просмотреть в `example.txt`.

#### Сборка контейнеров

Установка зависимостей:
```sh
make deps
```

Для сборки Docker-контейнеров выполните команду:
```sh
make build
```

#### Запуск контейнеров

Для запуска контейнеров выполните команды:
```sh
make up-app
make up-db
make up-broker
```

Если вам нужно запустить только определенные контейнеры:
- Запуск только приложения:
  ```sh
  make up-app
  ```
- Запуск только базы данных:
  ```sh
  make up-db
  ```
- Запуск только Kafka брокера:
  ```sh
  make up-broker
  ```
- Запуск Redis:
  ```sh
  make up-redis
  ```
- Запуск контейнеров для тестов:
  ```sh
  make up-db-test
  make up-broker-test
  make up-redis-test
  ```

#### Остановка контейнеров

Для остановки всех контейнеров выполните команду:
```sh
make down
```

Если вам нужно остановить только определенные контейнеры:
- Остановка только приложения:
  ```sh
  make down-app
  ```
- Остановка только базы данных:
  ```sh
  make down-db
  ```
- Остановка только Kafka брокера:
  ```sh
  make down-broker
  ```
- Остановка Redis:
  ```sh
  make down-redis
  ```
- Остановка контейнеров для тестов:
  ```sh
  make down-db-test
  make down-broker-test
  make down-redis-test
  ```

#### Миграция базы данных

Для применения миграций базы данных выполните команду:
```sh
make migrate
```

Для выполнения миграций в тестовой базе данных:
```sh
make migrate-test
```

Для отката миграций базы данных:
```sh
make demigrate
```

Для отката миграций в тестовой базе данных:
```sh
make demigrate-test
```

#### Основные процессы

- Запуск сервера:
  ```sh
  make server
  ```
- Запуск SQL-клиента:
  ```sh
  make sql
  ```
- Запуск SQL-клиента для тестов:
  ```sh
  make sql-test
  ```
- Запуск redis-cli:
  ```sh
  make redis-cli
  ```
- Запуск redis-cli для тестов:
  ```sh
  make redis-cli-test
  ```

#### Открытие shell внутри контейнеров

- Открытие shell внутри контейнера приложения:
  ```sh
  make shell-app
  ```
- Открытие shell внутри контейнера базы данных:
  ```sh
  make shell-db
  ```
- Открытие shell внутри контейнера тестовой базы данных:
  ```sh
  make shell-db-test
  ```
- Открытие shell внутри контейнера Kafka брокера:
  ```sh
  make shell-broker
  ```
- Открытие shell внутри контейнера тестового Kafka брокера:
  ```sh
  make shell-broker-test
  ```
- Открытие shell внутри контейнера Redis:
  ```sh
  make shell-redis
  ```
- Открытие shell внутри контейнера тестового Redis:
  ```sh
  make shell-redis-test
  ```

#### Работа с моками

Для создания моков выполните команду:
```sh
make mocks
```

#### Тестирование

Перед запуском тестов, не забудьте сгенерировать моки через `make mocks`, если вы ещё их не сгенерировали.

Для запуска всех тестов выполните команду:
```sh
make tests
```

Для запуска отдельных тестов:
- Юнит-тесты:
  ```sh
  make tests-unit
  ```
- Интеграционные тесты:
  ```sh
  make tests-int
  ```

#### Генерация gRPC файлов

Для генерации необходимых для работы с gRPC go-файлов выполните команду:
```sh
make generate
```

Если вам нужно сгенерировать только gRPC файлы без зависимостей, выполните команду:
```sh
make generate-no-deps
```
