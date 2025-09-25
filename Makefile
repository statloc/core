lint:
	go tool golangci-lint run

test:
	go test -v ./...

coverage:
	go test -coverprofile=.coverage ./...

clean:
	rm .coverage
