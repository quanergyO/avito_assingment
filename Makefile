.PHONY: run build compose-run compose-rebuild dbrun test

run:
	go run cmd/main.go

build:
	mkdir build || true
	go build -o ./build/server.out cmd/main.go

compose-run:
	docker-compose up

compose-rebuild:
	docker-compose up --build

dbrun:
	docker run --name=postgres -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres

migrateUp:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up

migrateDown:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down

teste2e:
	docker run --name=avito-test-bd -e POSTGRES_PASSWORD='qwerty' -p 5438:5432 -d --rm postgres
	go test -count=1 ./tests/ || true
	docker stop avito-test-bd

test:
	go test ./... -v