include ./dev/env.dev

DB_URL="postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

dependencies:
	go get
	go mod tidy
	go mod vendor

build:
	go build

run:
	go build
	go run main.go

setup-dev-up:
	docker-compose down
	docker-compose up -d

setup-dev-down:
	docker-compose down

cover:
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out -o cover.html
	google-chrome cover.html

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-up:
	migrate -path database/migrations -database "$(DB_URL)" -verbose up

migrate-up-1:
	migrate -path database/migrations -database "$(DB_URL)" -verbose up 1

migrate-down:
	migrate -path database/migrations -database "$(DB_URL)" -verbose down

migrate-down-1:
	migrate -path database/migrations -database "$(DB_URL)" -verbose down 1

swaggo-install:
	go install github.com/swaggo/swag/cmd/swag@latest

swaggo-generate:
	swag init

