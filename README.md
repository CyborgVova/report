# Report

### Простой RESTapi сервис для обработки отчетов

- Создание и запуск сервиса: `make` (требуется настроить данные базы данных в Dockerfile)

- Создание и запуск сервиса c подключенной тестовой базой: `make compose`

- Запросы:

  - `/api/v1/get_report?report_id=<number>` (1, 2, 3 ...)
  - `/api/v1/get_observation_time?model_id<number>` (10, 20, 30 ...)
  - `/api/v1/set_report` (body of request {"report_info":"\<character\>"})

- Тестовая база данных `./test/report.sql`
