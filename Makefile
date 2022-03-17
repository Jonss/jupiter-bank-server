
.PHONY: test run

test:
	go test -v --cover ./...

make build:
	CGO_ENABLED=0 GOOS=linux go build -o jupiterbank/bin cmd/jupiterbank/main.go

run:
	go run cmd/jupiterbank/main.go
