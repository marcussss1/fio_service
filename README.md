# Сервис информации о человеке(ФИО, возраст, пол, национальность)

## Запуск

Для запуска контейнера
```shell
make run
```

Для остановки контейнера
```shell
make stop
```

Для генерации api в формате swagger
```shell
make generate_api
```

## О проекте

Конфигурация в .env, api в формате swagger доступно по ручке "/docs/*". Пример тестов на хендлеры в internal/fio_service/delivery/http/update_people_test.go, на логику в internal/fio_service/usecase/create_people_test.go, на хранилище в internal/fio_service/repository/update_people_by_id_test.go. Логгирование осуществляется в internal/middleware/middleware.go.
