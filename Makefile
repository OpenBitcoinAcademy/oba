.PHONY: build test test-cover lint translate-check run clean

build:
	go build -o oba ./cmd/oba

test:
	go test ./...

test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	gofmt -l . | grep -v vendor | (! read) || (echo "gofmt found unformatted files" && exit 1)
	go vet ./...
	$(shell go env GOPATH)/bin/staticcheck ./...

translate-check:
	go run ./tools/translate-check/...

run: build
	./oba

clean:
	rm -f oba coverage.out coverage.html
