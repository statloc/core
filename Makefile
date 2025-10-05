.DEFAULT_GOAL := check

check: lint test

lint:
	go tool golangci-lint run

build:
	go mod download

test:
	go test ./...

coverage:
	go test -coverprofile=.coverage ./...

clean:
	rm -f .coverage
	go clean -testcache
