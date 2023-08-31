# [Тестовое задание](task/README.md) для стажёра Backend
# Сервис динамического сегментирования пользователей

### Swagger:
```
http://localhost/docs
```

### Для запуска приложения:

#### 1. Собрать проект:
```
make build
```

#### 2. Запустить проект:
```
make run
```

### Примеры запросов:
#### 1. /auth/register

***Запрос:***
```
curl -X 'POST' \
  'http://localhost:8000/auth/register' \
  -H 'accept: application/json' \
  -d ''
```
***Ответ:***
```
{
  "id": 1
}
```

#### 2. /api/segments

***Запрос:***
```
curl -X 'POST' \
  'http://localhost:8000/api/segments' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "slug": "AVITO_TEST_SEGMENT"
}'
```
***Ответ:***
```
{
  "id": 1
}
```

#### 3. /api/users/{id}/add_to_segment

***Запрос:***
```
curl -X 'POST' \
  'http://localhost:8000/api/users/1/add_to_segment' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[
  "AVITO_TEST_SEGMENT"
]'
```
***Ответ:***
```
{
  "status": "ok"
}
```

#### 4. /api/users/{id}/show_active_segments

***Запрос:***
```
curl -X 'GET' \
  'http://localhost:8000/api/users/1/show_active_segments' \
  -H 'accept: application/json'
```
***Ответ:***
```
{
  "User_id": 1,
  "Slugs": [
    "AVITO_TEST_SEGMENT"
  ]
}
```

#### 5. /api/users/{id}/delete_from_segment

***Запрос:***
```
curl -X 'DELETE' \
  'http://localhost:8000/api/users/1/delete_from_segment' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '[
  "AVITO_TEST_SEGMENT"
]'
```
***Ответ:***
```
{
  "status": "ok"
}
```

#### 6. /api/segments/{slug}

***Запрос:***
```
curl -X 'DELETE' \
  'http://localhost:8000/api/segments/AVITO_TEST_SEGMENT' \
  -H 'accept: application/json'
```
***Ответ:***
```
{
  "status": "ok"
}
```
