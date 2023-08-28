[![Golang](https://img.shields.io/badge/Go-v1.21-EEEEEE?logo=go&logoColor=white&labelColor=00ADD8)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

<div align="center">
    <h1>Сервис динамического сегментирования пользователей</h1>
    <h5>
        Микросервис, написанный на языке Golang, для добавления пользователей в определенные группы (сегменты),
удаления пользователей из сегментов с возможностью автоматизировать данные процессы. Также присутствует возможность
получить сводный отчет по всем операциям (с возможность указать конкретных пользователей) в формате csv-файла или ссылки на csv файл.
    </h5>
    <p>
        Russian | <a href="README.en.md">English</a> 
    </p>
</div>

### Используемые технологии:
- [Golang](https://go.dev), [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/), [Nginx](https://nginx.org/ru/)
- [REST](https://ru.wikipedia.org/wiki/REST), [Swagger UI](https://swagger.io/tools/swagger-ui/)
- [JWT Аутентификация](https://jwt.io/)

### Установка
`go get github.com/zenorachi/dynamic-user-segmentation`

### Начало работы
1. **Чтобы сервис корректно обрабатывал запросы на получение ссылки на csv-файл,
необходимо:**
    * Зарегистрировать приложение в [Google Cloud Platform](https://developers.google.com/workspace/guides/create-project);
    * Создать сервисный аккаунт и секретный ключ для него;
    * Добавить в директорию secrets полученный секретный ключ;
    * Изменить переменную окружения `GDRIVE_CREDENTIALS` в .env.
> **Подсказка:** сервис можно запустить без интеграции с Google Drive. В таком случае
> при запросе на получение ссылки на csv-файл будет ошибка, говорящая, что сервис GDrive недоступен.
2. **Настройка переменных окружения (создайте файл .env в корне проекта):**
```dotenv
# Database
export DB_HOST=
export DB_PORT=
export DB_USER=
export DB_NAME=
export DB_SSLMODE=
export DB_PASSWORD=

# Local database
export LOCAL_DB_PORT=

# Postgres service
export POSTGRES_PASSWORD=

# Password Hasher
export HASH_SALT=
export HASH_SECRET=

# Path to Google Drive credentials.json
export GDRIVE_CREDENTIALS=./secrets/your_credentials_file.json

# Gin mode (optional, default - release)
export GIN_MODE=
```
> **Подсказка:** если вы запускаете проект с помощью Docker, установите `DB_HOST`=postgres (как имя сервиса Postgres в docker-compose).
3. **Запуск сервиса:**
```shell
make
```
4. **Чтобы протестировать работу сервиса, можно перейти по адресу
`http://localhost:8080/docs/index.html` для получения Swagger документации.**

---

### [Примеры запросов]

---

### Дополнительные возможности:
1. **Запуск тестов:**
```shell
make test
```
2. **Запуск линтера:**
```shell
make lint
```
3. **Создание файлов миграций:**
```shell
make migrate-create
```
4. **Миграции вверх/вниз:**
```shell
make migrate-up
```
```shell
make migrate-down
```