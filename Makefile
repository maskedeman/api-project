run watch-mode:
	go mod tidy
	gow run cmd/api-project/main.go

run:
	go mod tidy
	go run cmd/api-project/main.go

tidy:
	go mod tidy

migrate:
	go run cmd/cli/migrations.go

lint:
	golangci-lint run
