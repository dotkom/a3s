<br />
<p align="center">
  <h3 align="center">Attendance as a Service</h3>

  <p align="center">
    Microservice for signing up users for events
    <br />
    <a href="https://github.com/dotkom/a3s/issues">Report Bug</a>
    |
    <a href="https://github.com/dotkom/a3s">Request Feature</a>
  </p>
</p>

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
- [Development](#development)

## Getting Started

The service is a backend GraphQL API which is written in Golang. It uses [gqlgen](gqlgen) to create strongly typed
GraphQL APIs and [Entgo](entgo) as a ORM in front of a PostGreSQL database.

### Prerequisites

To start development server you need the following components on your local machine:

1. Golang (1.17)
2. PostGreSQL (10.0+) (optionally use Docker + Docker Compose build)

### Development

To start a development server on your machine follow these steps:

1. Set local environment variables required to run the application:
    ```sh
    export POSTGRES_USER=admin
    export POSTGRES_PASSWORD=admin
    export POSTGRES_DB=a3s
    # Used by Migration script
    export POSTGRESQL_URL="postgres://$POSTGRES_USER:POSTGRES_PASSWORD@localhost:5432/$POSTGRES_DB?sslmode=disable"
    ```

2. Download Go dependencies with `go mod download`
3. Generate `gqlgen` code with `go run github.com/99designs/gqlgen generate`
4. Generate Ent entities with `go generate ./ent`
5. Migrate the Ent schema to your local PostGreSQL database with `go run bin/migrate.go`
6. Start the server with `go run bin/main.go`

[gqlgen]: https://github.com/99designs/gqlgen
[entgo]: https://entgo.io/