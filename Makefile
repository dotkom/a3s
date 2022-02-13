# Define the command line invocation of Go, if necessary:
ifeq ($(GO),)
    GO := go
endif

.PHONY: db download gen migrate main help

app: db download gen migrate run-main

run-main:
	@ $(GO) run cmd/main.go

db:
	@ docker-compose up -d psql

download:
	@ $(GO) mod download

gen:
	@ $(GO) generate ./ent && $(GO) run github.com/99designs/gqlgen generate

migrate:
	@ $(GO) run cmd/migrate.go
