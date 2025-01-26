build:
	@go build -o bin/af1-spider

run: build 
	@./bin/af1-spider

test:
	go test -v ./...