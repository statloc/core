.DEFAULT_GOAL := check

check: lint test

lint:
	go tool golangci-lint run

build:
	go mod download

test:
	go test -coverprofile=.coverage ./... | column -t

clean:
	rm -f .coverage
	go clean -testcache
