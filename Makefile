.SILENT:

build:
	docker-compose build

run:
	docker-compose up

swag:
	swag init -g cmd/main.go
