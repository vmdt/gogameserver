.PHONY: run swag

run:
	go run cmd/api/main.go

socket:
	cd socket-service && npm run start:dev

swag:
	swag init -g ./cmd/api/main.go