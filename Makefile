.PHONY: help run build test fmt vet up down logs

# Local-only convenience: .env is deliberately ignored by Git. Deployment
# environments provide these values through their platform's secret/config
# mechanism instead of a checked-in file.
ifneq (,$(wildcard ./.env))
include .env
export
endif

help:
	@echo "Available commands:"
	@echo "  make run    Start the API locally"
	@echo "  make build  Build the API binary"
	@echo "  make test   Run tests"
	@echo "  make fmt    Format Go code"
	@echo "  make vet    Run Go's static analysis"
	@echo "  make up     Start local infrastructure"
	@echo "  make down   Stop local infrastructure"
	@echo "  make logs   Follow local infrastructure logs"

run:
	go run ./cmd/api

build:
	go build -o bin/zipcrypt-api ./cmd/api

test:
	go test ./...

fmt:
	gofmt -w $$(find . -name '*.go' -not -path './vendor/*')

vet:
	go vet ./...

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f
