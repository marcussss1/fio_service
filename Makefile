.PHONY: run
run:
	docker compose -f docker-compose.yml up

.PHONY: stop
stop:
	docker compose -f docker-compose.yml down

.PHONY: mocks
mocks: |
	mockgen --source=internal/fio_service/usecase.go > internal/mocks/fio_service/usecase.go
	mockgen --source=internal/fio_service/repository.go > internal/mocks/fio_service/repository.go
	mockgen --source=internal/third_service/usecase.go > internal/mocks/third_service/usecase.go
