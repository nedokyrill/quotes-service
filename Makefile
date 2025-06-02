build-app:
	@go build -o bin/quotes-service cmd/app/main.go

run:build-app
	@./bin/quotes-service

new-migrate:
	@migrate create -ext sql -dir db/migrations -seq ${name}

ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

DATABASE_URL := $(POSTGRESQL_URL)

migrate-up:
	@migrate -database "$(DATABASE_URL)" -path db/migrations up

migrate-down:
	@migrate -database "$(DATABASE_URL)" -path db/migrations down 1