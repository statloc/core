.DEFAULT_GOAL := check

check: lint test

lint:
	go tool golangci-lint run

build:
	go mod download

test:
	go test ./... -coverprofile=.coverage -timeout=10s -race

clean:
	rm -f .coverage
	go clean -testcache
