
.PHONY: test run build env-up env-down new-migration 

test:
	go test -v --cover ./...

build:
	CGO_ENABLED=0 GOOS=linux go build -o jupiterbank/bin cmd/jupiterbank/main.go

run:
	go run cmd/jupiterbank/main.go

env-up:
	docker-compose up --build -d

env-down:
	docker-compose down --remove-orphans

new-migration:
	migrate create -ext sql -dir pkg/db/migrations -seq $(name)

cover:
	go tool cover -html=coverage.ou
