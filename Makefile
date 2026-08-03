.PHONY: run tidy fmt compose-up compose-down

run:
	go run cmd/main.go

tidy:
	go mod tidy

fmt:
	gofmt -w .

compose-up:
	docker compose -f compose.yaml up -d

compose-down:
	docker compose -f compose.yaml down
