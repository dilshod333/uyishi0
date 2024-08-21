migrate:
	migrate create -dir ./internal/db -ext sql db

migrate-up:
	migrate -path ./internal/db -database "postgres://postgres:dilshod@localhost:5432/t?sslmode=disable" up

migrate-down:
	migrate -path ./internal/db -database "postgres://postgres:dilshod@localhost:5432/t?sslmode=disable" down

migrate-force:
	migrate -path ./internal/db -database "postgres://postgres:dilshod@localhost:5432/t?sslmode=disable" force 20240701184310

run:
	@go run cmd/main.go 

tidy:
	@go mod tidy
