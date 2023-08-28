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
        Russian | <a href="README.md">English</a> 
    </p>
</div>

### Используемые технологии:
- [Golang](https://go.dev), [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/), [Nginx](https://nginx.org/ru/)
- [REST](https://ru.wikipedia.org/wiki/REST), [Swagger UI](https://swagger.io/tools/swagger-ui/)
- [JWT Аутентификация](https://jwt.io/)

### Установка
```shell
git clone git@github.com/zenorachi/dynamic-user-segmentation.git
```

### Начало работы
1. **Чтобы сервис корректно обрабатывал запросы на получение ссылки на csv-файл,
   необходимо:**
    * Зарегистрировать приложение в [Google Cloud](https://developers.google.com/workspace/guides/create-project);
    * Создать сервисный аккаунт и секретный ключ для него;
    * Добавить в директорию `secrets` полученный секретный ключ;
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
   http://localhost:8080/docs/index.html для получения Swagger документации.**

---

### [Примеры запросов](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md)

**[Пользователи](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#Пользователи)**
* [Регистрация](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#1-регистрация)
* [Аутентификация](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#2-аутентификация)
* [Обновление токена](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#3-обновление-токена)

**[Сегменты](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#Сегменты)**
* [Создание сегмента](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#1-создание-сегмента)
* [Создание сегмента с указанием процента автоматического добавления](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#2-создание-сегмента-с-указанием-процента-автоматического-добавления)
* [Удаление сегмента по имени](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#3-удаление-сегмента-по-имени)
* [Удаление сегмента по ID](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#4-удаление-сегмента-по-id)
* [Получение всех сегментов](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#5-получение-всех-сегментов)
* [Получение сегмента по ID](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#6-получение-сегмента-по-id)

**[Операции добавления / удаления сегментов пользователя](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#операции-добавления--удаления-сегментов-пользователя)**
* [Добавление сегментов пользователю по списку имен](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#1-добавление-сегментов-пользователю-по-списку-имен)
* [Добавление сегментов пользователю по списку имен с указанием времени жизни](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#2-добавление-сегментов-пользователю-по-списку-имен-с-указанием-времени-жизни)
* [Добавление сегментов пользователю по списку ID](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#3-добавление-сегментов-пользователю-по-списку-id)
* [Добавление сегментов пользователю по списку ID с указанием времени жизни](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#4-добавление-сегментов-пользователю-по-списку-id-с-указанием-времени-жизни)
* [Удаление сегментов у пользователя по списку имен](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#5-удаление-сегментов-у-пользователя-по-списку-имен)
* [Удаление сегментов у пользователя по списку ID](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#6-удаление-сегментов-у-пользователя-по-списку-id)

**[Отношения пользователи-сегменты](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#отношения-пользователи-сегменты)**
* [Получение активных сегментов пользователя](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#1-получение-активных-сегментов-пользователя)
* [Получение активных пользователей сегмента](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#2-получение-активных-пользователей-сегмента)

**[Отчеты](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#Отчеты)**
* [Получение истории операций](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#1-получение-истории-операций)
* [Получение истории операций в виде csv-файла](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#2-получение-истории-операций-в-виде-csv-файла)
* [Получение истории операций в виде ссылки на csv-файл](https://github.com/zenorachi/dynamic-user-segmentation/blob/main/docs/examples/01-requests.ru.md#3-получение-истории-операций-в-виде-ссылки-на-csv-файл)

---

### Дополнительные возможности
1. **Запуск тестов**
```shell
make test
```
2. **Запуск линтера**
```shell
make lint
```
3. **Создание файлов миграций**
```shell
make migrate-create
```
4. **Миграции вверх/вниз**
```shell
make migrate-up
```
```shell
make migrate-down
```