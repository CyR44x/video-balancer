# Video Balancer

Балансировщик видео-трафика принимает gRPC запросы с URL видео и выполняет редирект:
- Каждый 10-й запрос возвращает оригинальный URL.
- Остальные запросы перенаправляются на CDN, используя переменную окружения `CDN_HOST`.


## Сборка и запуск

### Через Docker Compose

   ```bash
   docker-compose up --build
   ```

### Локально
1. Установить переменную окружения:
   ```bash
   export CDN_HOST=cdn.example.com
   ```
2. Запустить сервер:
   ```bash
   go run ./cmd/server
   ```
Проверить работу сервиса через [grpcurl](https://github.com/fullstorydev/grpcurl
):
   ```bash
   grpcurl -plaintext -d '{"video": "http://s1.origin-cluster/video/123/xcg2djHckad.m3u8"}' localhost:50051 my.custom.server.Service/Method
   ```

## Нагрузочное тестирование

Для проверки производительности используйте [ghz](https://github.com/bojand/ghz):
   ```bash
   ghz --insecure \
    --proto proto/service.proto \
    --call my.custom.server.Service/Method \
    -d '{"video": "http://s1.origin-cluster/video/123/xcg2djHckad.m3u8"}' \
    localhost:50051 -n 100000 -c 100
   ```
Параметры:

- n – общее количество запросов
- c – число параллельных соединений

### Результаты нагрузочного тестирования
```
Summary:
Count:        10000
Total:        857.53 ms
Slowest:      20.00 ms
Fastest:      0 ns
Average:      2.63 ms
Requests/sec: 11661.42

Response time histogram:
0.000  [122]  |∎
2.000  [3974] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
4.000  [4419] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
6.000  [1195] |∎∎∎∎∎∎∎∎∎∎∎
8.000  [174]  |∎∎
10.000 [46]   |
12.000 [19]   |
14.001 [1]    |
16.001 [28]   |
18.001 [18]   |
20.001 [4]    |

Latency distribution:
10 % in 1.00 ms
25 % in 1.62 ms
50 % in 2.17 ms
75 % in 3.35 ms
90 % in 4.39 ms
95 % in 5.02 ms
99 % in 8.33 ms

Status code distribution:
[OK]   10000 responses
```
## Идеи для доработки:

- Добавить тесты
- Логирование и мониторинг
- Валидация запроса
- Аутентификация и авторизация

