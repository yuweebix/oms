## «Утилита для управления ПВЗ»
### Список команд

1. **Принять заказ от курьера**
   На вход принимается ID заказа, ID получателя и срок хранения. Заказ нельзя принять дважды. Если срок хранения в прошлом, приложение должно выдать ошибку. Список принятых заказов необходимо сохранять в файл. Формат файла остается на выбор автора.
2. **Вернуть заказ курьеру**
   На вход принимается ID заказа. Метод должен удалять заказ из вашего файла. Можно вернуть только те заказы, у которых вышел срок хранения и если заказы не были выданы клиенту.
3. **Выдать заказ клиенту**
   На вход принимается список ID заказов. Можно выдавать только те заказы, которые были приняты от курьера и чей срок хранения меньше текущей даты. Все ID заказов должны принадлежать только одному клиенту.
4. **Получить список заказов**
   На вход принимается ID пользователя как обязательный параметр и опциональные параметры.
   Параметры позволяют получать только последние N заказов или заказы клиента, находящиеся в нашем ПВЗ.
5. **Принять возврат от клиента**
   На вход принимается ID пользователя и ID заказа. Заказ может быть возвращен в течение двух дней с момента выдачи. Также необходимо проверить, что заказ выдавался с нашего ПВЗ.
6. **Получить список возвратов**
   Метод должен выдавать список пагинированно.

### Описание архитектуры

![alt text](docs/UML-Class-Diagram.png)
в формате [UML Class Diagram](https://www.drawio.com/blog/uml-class-diagrams)

### Пример использования утилиты

Утилита содержит перечень команд, что можно производить над заказами и возвратами.

#### Заказы
  - `orders accept  [flags]`: Принять заказ от курьера
  - `orders return  [flags]`: Вернуть заказ курьеру
  - `orders deliver [flags]`: Выдать заказ клиенту
  - `orders list    [flags]`: Получить список заказов
	
#### Возвраты
  - `returns accept [flags]`: Принять возврат от клиента
  - `returns list   [flags]`: Получить список возвратов

### Руководство по работе с проектом

#### Настройка окружения

Создайте файл `.env` в корне проекта и добавьте необходимые переменные окружения. Пример файла `.env` можно найти среди файлов проекта.

Соберите Docker-контейнеры, создайте соответствующую базу данных и выполните миграции. Всё это можно сделать с помощью команд, описанных ниже.

#### Сборка контейнеров

Для сборки Docker-контейнеров выполните команду:
```sh
make build
```

Вы также можете собрать отдельные контейнеры:
- Для сборки только приложения:
  ```sh
  make build-app
  ```
- Для сборки контейнера тестов:
  ```sh
  make build-test
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
- Запуск контейнеров для тестов:
  ```sh
  make up-app-test
  make up-db-test
  make up-broker-test
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
- Остановка контейнеров для тестов:
  ```sh
  make down-app-test
  make down-db-test
  make down-broker-test
  ```

#### Создание баз данных

Для создания баз данных выполните команды:
```sh
make create-db
make create-db-test
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

- Запуск утилиты:
  ```sh
  make cli
  ```
- Запуск утилиты тестов:
  ```sh
  make cli-test
  ```
- Запуск SQL-клиента:
  ```sh
  make sql
  ```
- Запуск SQL-клиента для тестов:
  ```sh
  make sql-test
  ```
- Просмотр логов приложения:
  ```sh
  make log
  ```
- Просмотр логов тестов:
  ```sh
  make log-test
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
- Открытие shell внутри контейнера тестов:
  ```sh
  make shell-app-test
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

#### Работа с моками

Для создания моков выполните команды:
- Утилиты:
  ```sh
  make mocks-cli
  ```
- Домена:
  ```sh
  make mocks-domain
  ```

#### Тестирование

Прежде чем проводить тестирование, нужно запустить контейнер для тестов. Тесты будут проводиться внутри контейнера.

Для запуска всех тестов выполните команду:
```sh
make test-all
```

Для запуска отдельных тестов:
- Тесты утилиты:
  ```sh
  make test-cli
  ```
- Тесты домена:
  ```sh
  make test-domain
  ```
- Тесты репозитория:
  ```sh
  make test-repository
  ```
- Тесты Kafka:
  ```sh
  make test-kafka
  ```

Если вы хотите провести тесты локально, не забудьте запустить тестовую базу данных. И в случае кафки - тестовый брокер.