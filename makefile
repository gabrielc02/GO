.DEFAULT_GOAL := build

# Formating
fmt:
	go fmt ./...
.PHONY:fmt

# Lint
lint: fmt
	golint ./...
.PHONY:lint

# Vetting
vet: fmt
	go vet ./...
	shadow ./...
.PHONY:vet

# BUILD
build: vet
	go build ./*.go
/PHONY: build

