test:
	go test -race ./...

lint:
	docker run --rm -it -v $(PWD):/app -w /app golangci/golangci-lint:latest golangci-lint run -v