build:
	@go build -o bin/gobank

run: build
	@./bin/gobank

dev:
	@which gin > /dev/null || go install github.com/codegangsta/gin@latest
	@gin --bin bin/gobank run main.go

test:
	@go test -v ./...
