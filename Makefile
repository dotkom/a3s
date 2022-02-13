# Define the command line invocation of Go, if necessary:
ifeq ($(GO),)
    GO := go
endif

.PHONY: db download gen migrate main help

# Starts database, download dependencies, generate GraphQL and database schema, add Ent schema to database, and start application.
app: db download gen migrate main

main:
	@ $(GO) run cmd/main.go

db:
	@ docker-compose up -d psql

download:
	@ $(GO) mod download

gen:
	@ $(GO) generate ./ent && $(GO) run github.com/99designs/gqlgen generate

migrate:
	@ $(GO) run cmd/migrate.go

help:
	@ echo "Usage	:  make <target>"
	@ echo "Targets	:"
	@ echo "	app - Execute all targets below"
	@ echo "	db - Start database"
	@ echo "	download - Download dependencies"
	@ echo "	gen - Generate GraphQL and database schema"
	@ echo "	migrate - Add Ent schema to database"
	@ echo "	main - Start application"
