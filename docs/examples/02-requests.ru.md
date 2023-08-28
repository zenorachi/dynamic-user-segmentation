# Примеры запросов

---

## Пользователи
### 1. Регистрация
* Пример запроса:
```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/users/sign-up' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "email": "maksim-go@gmail.com",
  "login": "maksim-go",
  "password": "qwerty123"
}'
```
* Пример ответа:
```json
{
  "id": 1
}
```

### 2. Аутентификация
* Пример запроса:
```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/users/sign-in' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "login": "maksim-go",
  "password": "qwerty123"
}'
```
* Пример ответа:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMyMzA4NDMsInN1YiI6IjEifQ.wJAI1tRNd3NFQ-KYw5e3Iy8RuXHRqIJoqeTAwXdMbNc"
}
```

### 3. Обновление токена
* Пример запроса:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/users/refresh' \
  -H 'accept: application/json'
```
* Пример ответа:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMyMzA4ODAsInN1YiI6IjEifQ.abo2LM6Xu13IHueO3Ia8K1UJh966S89QOifQ4cpPwqU"
}
```

---

## Сегменты
### 1. Создание сегмента
* Пример запроса:
```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/segments/create' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU4MjE3NjEsInN1YiI6IjEifQ.DSf-xaovGIiwmRKB3Zfa4E1eFvwDvYPv3RciLOEIvU0' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "AVITO-INTERNS"
}'
```
* Пример ответа:
```json
{
  "id": 1
}
```
### 2. Создание сегмента с указанием процента автоматического добавления
* Пример запроса:
```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/segments/create' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMyMzU5MDIsInN1YiI6IjEifQ.DvE1rs1OVyocP669GtCmD2Lk7DwbH37jsRePhfCP9Nk' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "VOICE-MESSAGE",
  "assign_percent": 50
}'
```
* Пример ответа:
```json
{
  "id": 1
}
```

### 3. Удаление сегмента по имени
* Пример запроса:
```shell
curl -X 'DELETE' \
  'http://localhost:8080/api/v1/segments/delete/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU4MjE3NjEsInN1YiI6IjEifQ.DSf-xaovGIiwmRKB3Zfa4E1eFvwDvYPv3RciLOEIvU0' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "TEST-PERCENT-10"
}'
```
* Пример ответа отсутствует
> **Пояснение:** если удаление прошло успешно, сервер вернет 204 код (NO CONTENT).

### 4. Удаление сегмента по ID
* Пример запроса:
```shell
curl -X 'DELETE' \
  'http://localhost:8080/api/v1/segments/delete_by_id/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU4MjE3NjEsInN1YiI6IjEifQ.DSf-xaovGIiwmRKB3Zfa4E1eFvwDvYPv3RciLOEIvU0' \
  -H 'Content-Type: application/json' \
  -d '{
  "id": 8
}'
```
* Пример ответа отсутствует
> **Пояснение:** если удаление прошло успешно, сервер вернет 204 код (NO CONTENT).

### 5. Получение всех сегментов
* Пример запроса:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/segments/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMyMzU5MDIsInN1YiI6IjEifQ.DvE1rs1OVyocP669GtCmD2Lk7DwbH37jsRePhfCP9Nk'
```
* Пример ответа:
```json
{
  "segments": [
    {
      "id": 7,
      "name": "AVITO-INTERNS"
    },
    {
      "id": 10,
      "name": "VOICE-MESSAGE",
      "assign_percent": 0.5
    },
    {
      "id": 11,
      "name": "DISCOUNT-30",
      "assign_percent": 0.5
    }
  ]
}
```

### 6. Получение сегмента по ID
>**Подсказка:** ID сегмента указывается параметром в URL.
* Пример запроса:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/segments/10' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU4Mjg0ODYsInN1YiI6IjEifQ.n1r4juv8EIvFdZr89Hb_9DRTa5vp2csaITWx-P3Le30'
```
* Пример ответа:
```json
{
  "segment": {
    "id": 10,
    "name": "VOICE-MESSAGE",
    "assign_percent": 0.5
  }
}
```

---

## Операции добавления / удаления сегментов пользователя
### 1. Добавление сегментов пользователю по списку имен
* Пример запроса:
```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/operations/add_segments_by_names/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU4Mjk2OTgsInN1YiI6IjEifQ.0n5dMO_JwxKpaesCTP-cxSwQ_PInydRqSAhBGDebhDA' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_id": 1,
  "segment_names": [
    "AVITO-INTERNS",
    "DISCOUNT-30"
  ]
}'
```
* Пример ответа:
```json
{
  "operation_ids": [
    82,
    83
  ]
}
```
### 2. Добавление сегментов пользователю по списку имен с указанием времени жизни
* Пример запроса:
```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/operations/add_segments_by_names/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU4Mjk2OTgsInN1YiI6IjEifQ.0n5dMO_JwxKpaesCTP-cxSwQ_PInydRqSAhBGDebhDA' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_id": 1,
  "segment_names": [
    "AVITO-INTERNS",
    "DISCOUNT-30"
  ],
  "ttl": "1h"
}'
```
* Пример ответа:
```json
{
  "operation_ids": [
    82,
    83
  ]
}
```

### 3. Добавление сегментов пользователю по списку ID
* Пример запроса:
```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/operations/add_segments_by_names/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU4Mjk2OTgsInN1YiI6IjEifQ.0n5dMO_JwxKpaesCTP-cxSwQ_PInydRqSAhBGDebhDA' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_id": 1,
  "segment_ids": [
    1,
    2
  ]
}'
```
* Пример ответа:
```json
{
  "operation_ids": [
    1,
    2
  ]
}
```
### 4. Добавление сегментов пользователю по списку ID с указанием времени жизни
* Пример запроса:
```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/operations/add_segments_by_names/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU4Mjk2OTgsInN1YiI6IjEifQ.0n5dMO_JwxKpaesCTP-cxSwQ_PInydRqSAhBGDebhDA' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_id": 1,
  "segment_ids": [
    1,
    2
  ],
  "ttl": "1h"
}'
```
* Пример ответа:
```json
{
  "operation_ids": [
    3,
    4
  ]
}
```

### 5. Удаление сегментов у пользователя по списку имен
* Пример запроса:
```shell
curl -X 'DELETE' \
  'http://localhost:8080/api/v1/operations/delete_segments_by_names/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU4MzExMTQsInN1YiI6IjEifQ.Cnnk6GRGFbBcSv5kDOtLfVgE2L9vubHpcxw5urnC_0A' \
  -H 'Content-Type: application/json' \
  -d '{
  "segment_names": [
    "VOICE-MESSAGE",
    "DISCOUNT-30"
  ],
  "user_id": 6
}'
```
* Пример ответа:
```json
{
  "operation_ids": [
    107,
    108
  ]
}
```
### 6. Удаление сегментов у пользователя по списку ID
* Пример запроса:
```shell
curl -X 'DELETE' \
  'http://localhost:8080/api/v1/operations/delete_segments_by_names/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU4MzExMTQsInN1YiI6IjEifQ.Cnnk6GRGFbBcSv5kDOtLfVgE2L9vubHpcxw5urnC_0A' \
  -H 'Content-Type: application/json' \
  -d '{
  "segment_ids": [
    1,
    2
  ],
  "user_id": 6
}'
```
* Пример ответа:
```json
{
  "operation_ids": [
    108,
    109
  ]
}
```

---

## Отношения пользователи-сегменты
### 1. Получение активных сегментов пользователя
>**Подсказка:** ID пользователя указывается параметром в URL.
* Пример запроса:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/users/active_segments/1' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU4MzExMTQsInN1YiI6IjEifQ.Cnnk6GRGFbBcSv5kDOtLfVgE2L9vubHpcxw5urnC_0A'
```
* Пример ответа:
```json
{
  "segments": [
    {
      "name": "AVITO-INTERNS"
    }
  ]
}
```

### 2. Получение активных пользователей сегмента
* Пример запроса:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/segments/active_users/7' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU4MzExMTQsInN1YiI6IjEifQ.Cnnk6GRGFbBcSv5kDOtLfVgE2L9vubHpcxw5urnC_0A'
```
* Пример ответа:
```json
{
  "users": [
    {
      "id": 1,
      "login": "maksim-go",
      "registered_at": "2023-08-28T10:24:17.927351Z"
    }
  ]
}
```

---

## Отчеты
### 1. Получение истории операций
>**Подсказка:** В запросе есть возможность указать размер страницы ("page_size": `size`) (сколько операций будет отображаться), а также не указывать user_ids (в таком случае вернется история по всем операциям для всех пользователей).
* Пример запроса:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/operations/history' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU4MzExMTQsInN1YiI6IjEifQ.Cnnk6GRGFbBcSv5kDOtLfVgE2L9vubHpcxw5urnC_0A' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_ids": [
    6
  ],
  "year": 2023,
  "month": 8
}'
```
* Пример ответа:
```json
{
  "operations": [
    {
      "user_id": 6,
      "segment_name": "TEST-PERCENT",
      "type": "added",
      "date": "2023-08-28T10:29:19.41186Z"
    },
    {
      "user_id": 6,
      "segment_name": "TEST-PERCENT-50",
      "type": "added",
      "date": "2023-08-28T10:29:33.986035Z"
    },
    {
      "user_id": 6,
      "segment_name": "TEST-PERCENT-50",
      "type": "deleted",
      "date": "2023-08-28T10:35:17.519398Z"
    },
    {
      "user_id": 6,
      "segment_name": "TEST-PERCENT",
      "type": "deleted",
      "date": "2023-08-28T10:35:50.447505Z"
    }
  ]
}
```

### 2. Получение истории операций в виде csv-файла
>**Подсказка:** В запросе есть также не указывать user_ids (в таком случае вернется файл со всеми операциями для всех пользователей).
* Пример запроса:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/reports/file' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU4MzExMTQsInN1YiI6IjEifQ.Cnnk6GRGFbBcSv5kDOtLfVgE2L9vubHpcxw5urnC_0A' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_ids": [
    1
  ],
  "year": 2023,
  "month": 8
}'
```
* Пример ответа:
```text
user-id,segment-name,type,date
1,TEST-PERCENT,added,2023-08-28 10:29:19
1,TEST-PERCENT,deleted,2023-08-28 10:35:50
1,AVITO-INTERN,added,2023-08-28 10:37:26
1,TEST-PERCENT-10,added,2023-08-28 10:37:26
1,AVITO-INTERN,added,2023-08-28 10:40:56
1,AVITO-THE-BEST,added,2023-08-28 10:40:56
```

### 3. Получение истории операций в виде ссылки на csv-файл
>**Подсказка:** В запросе есть также не указывать user_ids (в таком случае вернется файл со всеми операциями для всех пользователей).
* Пример запроса:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/reports/link' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTU4MzExMTQsInN1YiI6IjEifQ.Cnnk6GRGFbBcSv5kDOtLfVgE2L9vubHpcxw5urnC_0A' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_ids": [
    1
  ],
  "year": 2023,
  "month": 8
}'
```
* Пример ответа:
```json
{
  "link": "https://drive.google.com/file/d/1rBU6b17M_Edi9bqygu9sPU0Ve3IKPb1b/view?usp=sharing"
}
```