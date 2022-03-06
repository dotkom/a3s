# Define the command line invocation of Go, if necessary:
ifeq ($(GO),)
    GO := go
endif

.PHONY: db download gen migrate main help entity test

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

build: download gen
	@ $(GO) build -o bin/main -v cmd/main.go
	@ $(GO) build -o bin/migrate -v cmd/migrate.go

test: db download gen migrate
	@ $(GO) test -v ./repository/... ./resolvers/... ./graph/... ./ent/...

entity:
	@ sh ./generate-entity.sh $(name)

help:
	@ echo "Usage	:  make <target>"
	@ echo "Targets	:"
	@ echo "	app - Execute all targets below"
	@ echo "	db - Start database"
	@ echo "	download - Download dependencies"
	@ echo "	gen - Generate GraphQL and database schema"
	@ echo "	migrate - Add Ent schema to database"
	@ echo "	main - Start application"
<<<<<<< HEAD
	@ echo "	entity - Creates an Ent entity. Supply the entity name in PascalCase with the 'name' argument."
=======
	@ echo "	test - Test application"
>>>>>>> main
