start:
	@clear
	@go build -o bin/inline cmd/main.go
	@./bin/inline

tidy:
	@go mod tidy

test:
	@go test -v ./...