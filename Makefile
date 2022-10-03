.PHONY: default
default: build

.PHONY: build
build:
	GOOS=js GOARCH=wasm go build -o ./assets/wasm/polyglot.wasm

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run
