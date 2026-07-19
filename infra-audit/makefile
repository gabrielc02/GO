.DEFAULT_GOAL := build

# Formatação
fmt:
	go fmt ./...
.PHONY: fmt

# Lint
lint: fmt
	golint ./...
.PHONY: lint

# Vetting
vet: fmt
	go vet ./...
	go vet -vettool=/home/user/go/bin/shadow ./...
.PHONY: vet

# BUILD
build: vet
	go build -o infra-audit-program ./cmd/infra-audit
.PHONY: build
