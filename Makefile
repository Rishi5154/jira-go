run: build
	@./bin/jira

build:
	@go build -o bin/jira

test:
	@go test -v ./...