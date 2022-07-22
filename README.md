# Social Club

Тестовое задание


## Задача
Необходимо реализовать HTTP REST сервис, который должен предоставлять
API со следующими методами:
* Добавление факта коммуникации;
* Получение графа социальных связей, при этом граф должен содержать не только структуру по узлам и ребрам, но и иметь раздел со сводной информацией по количествам коммуникай между пользователями (минимальное, максимальное, среднее).

## Отправка сообщения
Для того, чтобы зарегистрировать факт коммуникации между пользователями, необходимо отправить Json-объект следующего формата на URL "/msg" с помощью метода POST:
```yaml
{
  "first_user" : id, 
  "second_user" : id
}
```
где id - уникальный неотрицательный идентификатор пользователя.

Очередность введения идентификаторов пользователя ни на что не влияет, так как учитывается факт взаимодействия.


## Получение информации о коммуникациях 
Для получения информации о коммуникациях необходимо сделать запрос на URL "/info" методом GET. Пример получаемого ответа:
```yaml
{
    "graph": [
        {
            "user_id": 1,
            "messages": "{User: 2; Send messages: 3}; {User: 5; Send messages: 2}; "
        },
        {
            "user_id": 2,
            "messages": "{User: 1; Send messages: 3}; {User: 3; Send messages: 2}; {User: 5; Send messages: 1}; "
        },
        {
            "user_id": 3,
            "messages": "{User: 2; Send messages: 2}; {User: 5; Send messages: 1}; "
        },
        {
            "user_id": 5,
            "messages": "{User: 2; Send messages: 1}; {User: 3; Send messages: 1}; {User: 1; Send messages: 2}; {User: 4; Send messages: 2}; "
        },
        {
            "user_id": 4,
            "messages": "{User: 5; Send messages: 2}; "
        }
    ],
    "min_value": 1,
    "avg_value": 2.2,
    "max_value": 3
}
```
## Запуск API
Для запуска API с in-memory хранилищем:
```
docker build -t <name> -f Dockerfile-memory .
docker run -p 5000:5000 -t <name>
```

Для запуска API с PostgreSQL:
```
docker-compose build --no-cache
docker-compose up
```