build:
	@go build -o bin/apigin

run: build
	@./bin/apigin 

test:
	@go test -v ./...

start:
	@docker compose up -d

stop:
	@docker compose down
