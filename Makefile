
.PHONY: test run

test:
	go test -v --cover ./...

run:
	go run cmd/jupiterbank/main.go
