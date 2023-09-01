lint:
	@echo "--> Running linter"
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.1
	@golangci-lint run --timeout=10m

build:
	@echo "--> Building app"
	go build -o Re_client cmd/main.go

install: build
	@echo "--> Installing app"
	mv Re_client /usr/local/bin

.PHONY: lint build install
