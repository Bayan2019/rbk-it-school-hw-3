# Домашнее задание 3

## REST API

### Добавить город пользователю

```bash
curl -X POST http://localhost:8080/api/v1/users/1/cities \
  -H "Content-Type: application/json" \
  -d '{
    "city": "Almaty"
  }'
```

### Список городов пользователя

```bash
curl -X GET http://localhost:8080/api/v1/users/1/cities
```

### Удалить город

```bash
curl -X DELETE http://localhost:8080/api/v1/users/1/cities/3
```

### Получить погоду по всем городам пользователя

```bash
curl -X GET http://localhost:8080/api/v1/users/1/weather
```

### Получить историю с фильтрацией

```bash
curl -X GET http://localhost:8080/api/v1/users/1/weather/history?city=astana"
```