.PHONY: build test test-cover lint translate-check run clean build-android

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

build-android:
	@echo "Requires: ANDROID_SDK_ROOT, ANDROID_NDK_HOME, gogio"
	@echo "Install gogio: go install gioui.org/cmd/gogio@latest"
	gogio -target android -name ObA -appid org.openbitcoinacademy.oba -o oba.apk ./cmd/oba

clean:
	rm -f oba oba.apk coverage.out coverage.html
