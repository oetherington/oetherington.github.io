all: run

test:
	go test -v -coverprofile=coverage.out

lint:
	golangci-lint run ./...

fmt:
	go fmt .

check-fmt:
	if [ "$$(gofmt -s -l . | wc -l)" -gt 0 ]; then exit 1; fi

run:
	go run .

build:
	go build

run-built:
	./oetherington.github.io

docs:
	godoc -http=127.0.0.1:6060

serve:
	caddy run

check: test lint check-fmt
