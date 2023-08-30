# Examples of requests

---

## Users

### 1. Registration
* Request example:
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
* Response example:
```json
{
  "id": 1
}
```

### 2. Authentication
* Request example:
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
* Response example:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw"
}
```

### 3. Refresh token
* Request example:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/users/refresh' \
  -H 'accept: application/json'
```
* Response example:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4NTIsInN1YiI6IjE0In0.cmXwc15TmNSI2mILSZjoqRhhtUN2AYZQu5had9OW07k"
}
```

---

## Segments

### 1. Create a segment
* Request example:
```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/segments/create' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "AVITO-INTERNS"
}'
```
* Response example:
```json
{
  "id": 1
}
```
### 2. Create a segment with an indication of the percentage of automatic addition
* Request example:
```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/segments/create' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "VOICE-MESSAGE",
  "assign_percent": 50
}'
```
* Response example:
```json
{
  "id": 1
}
```

### 3. Delete a segment by name
* Request example:
```shell
curl -X 'DELETE' \
  'http://localhost:8080/api/v1/segments/delete/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "TEST-PERCENT-10"
}'
```
* Response example отсутствует
> **Hint:** if the deletion was successful, the server will return code 204 (NO CONTENT).

### 4. Delete a segment by ID
* Request example:
```shell
curl -X 'DELETE' \
  'http://localhost:8080/api/v1/segments/delete_by_id/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw' \
  -H 'Content-Type: application/json' \
  -d '{
  "id": 8
}'
```
* Response example отсутствует
> **Hint:** if the deletion was successful, the server will return code 204 (NO CONTENT).

### 5. Get all segments
* Request example:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/segments/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw'
```
* Response example:
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

### 6. Get a segment by ID
>**Hint:** the segment ID is specified as a parameter in the URL.
* Request example:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/segments/10' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw'
```
* Response example:
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

## User segment addition / removal operations
### 1. Add segments to a user by a list of names
* Request example:
```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/operations/add_segments_by_names/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_id": 1,
  "segment_names": [
    "AVITO-INTERNS",
    "DISCOUNT-30"
  ]
}'
```
* Response example:
```json
{
  "operation_ids": [
    82,
    83
  ]
}
```
### 2. Add segments to the user by a list of names with an indication of the ttl
* Request example:
```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/operations/add_segments_by_names/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw' \
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
* Response example:
```json
{
  "operation_ids": [
    82,
    83
  ]
}
```

### 3. Add segments to a user by a list of ID
* Request example:
```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/operations/add_segments_by_names/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_id": 1,
  "segment_ids": [
    1,
    2
  ]
}'
```
* Response example:
```json
{
  "operation_ids": [
    1,
    2
  ]
}
```
### 4. Add segments to the user by a list of ID with an indication of the ttl
* Request example:
```shell
curl -X 'POST' \
  'http://localhost:8080/api/v1/operations/add_segments_by_names/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw' \
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
* Response example:
```json
{
  "operation_ids": [
    3,
    4
  ]
}
```

### 5. Delete segments from a user by a list of names
* Request example:
```shell
curl -X 'DELETE' \
  'http://localhost:8080/api/v1/operations/delete_segments_by_names/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw' \
  -H 'Content-Type: application/json' \
  -d '{
  "segment_names": [
    "VOICE-MESSAGE",
    "DISCOUNT-30"
  ],
  "user_id": 6
}'
```
* Response example:
```json
{
  "operation_ids": [
    107,
    108
  ]
}
```
### 6. Delete segments from a user by a list of ID
* Request example:
```shell
curl -X 'DELETE' \
  'http://localhost:8080/api/v1/operations/delete_segments_by_names/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw' \
  -H 'Content-Type: application/json' \
  -d '{
  "segment_ids": [
    1,
    2
  ],
  "user_id": 6
}'
```
* Response example:
```json
{
  "operation_ids": [
    108,
    109
  ]
}
```

---

## User-segment relations
### 1. Get active segments of a user
* Request example:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/users/active_segments/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMzQyNDUsInN1YiI6IjExIn0.XxlMAboSUE2Wey8wsbT4IxqmAXj6MfJfL7L8Pd3QthA' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_id": 1
}'
```
* Response example:
```json
{
  "segments": [
    {
      "name": "AVITO-INTERNS"
    }
  ]
}
```

### 2. Get active users of a segment
* Request example:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/segments/active_users/' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw' \
  -H 'Content-Type: application/json' \
  -d '{
  "segment_id": 1
}'
```
* Response example:
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

## Reports
### 1. Get operation history
>**Hint:** in the request, it is possible to specify the page size ("page_size": `size`) (how many operations will be displayed), and also not to specify user_ids (in this case, the history of all operations for all users will be returned).
* Request example:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/operations/history' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_ids": [
    6
  ],
  "year": 2023,
  "month": 8
}'
```
* Response example:
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

### 2. Get operation history in the form of a CSV file
>**Подсказка:** there is also no need to specify user_ids in the request (in this case, a file with all operations for all users will be returned).
* Request example:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/reports/file' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_ids": [
    1
  ],
  "year": 2023,
  "month": 8
}'
```
* Response example:
```text
user-id,segment-name,type,date
1,TEST-PERCENT,added,2023-08-28 10:29:19
1,TEST-PERCENT,deleted,2023-08-28 10:35:50
1,AVITO-INTERN,added,2023-08-28 10:37:26
1,TEST-PERCENT-10,added,2023-08-28 10:37:26
1,AVITO-INTERN,added,2023-08-28 10:40:56
1,AVITO-THE-BEST,added,2023-08-28 10:40:56
```

### 3. Get operation history in the form of a CSV file link
>**Hint:** there is also no need to specify user_ids in the request (in this case, a link to the file with all operations for all users will be returned).
* Request example:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/reports/link' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTMzMDc4MjksInN1YiI6IjE0In0.N1QBZb1uVZQGJ7vROHhCdlaySu1o31yjTzFLnVk_XYw' \
  -H 'Content-Type: application/json' \
  -d '{
  "user_ids": [
    1
  ],
  "year": 2023,
  "month": 8
}'
```
* Response example:
```json
{
  "link": "https://drive.google.com/file/d/1rBU6b17M_Edi9bqygu9sPU0Ve3IKPb1b/view?usp=sharing"
}
```