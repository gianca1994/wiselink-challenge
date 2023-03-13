SHELL:=/bin/bash -O extglob
BINARY=wiselink
VERSION=0.1.0
LDFLAGS=-ldflags "-X main.Version=${VERSION}"

#go tool commands
build:
	go build ${LDFLAGS} -o ${BINARY} src/main.go

run:
	@go run src/main.go

## docker compose
up:
	docker-compose up --build
down:
	docker-compose down --remove-orphans