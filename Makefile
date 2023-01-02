build:
	@go build -o bin/apigin

run: build
	@./bin/apigin 

test:
	@go test -v ./...
