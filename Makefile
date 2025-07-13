.PHONY: run swag

run:
	go run cmd/api/main.go

swag:
	swag init -g ./cmd/api/main.go