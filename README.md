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

The service is a backend GraphQL API which is written in Golang. It uses [gqlgen][gqlgen] to create strongly typed
GraphQL APIs and [Entgo][entgo] as a ORM in front of a PostGreSQL database.

### Prerequisites

To start development server you need the following components on your local machine:

1. Golang (1.17)
2. PostGreSQL (10.0+) (optionally use Docker + Docker Compose build)

### Development

#### With `make`

Make sure you have Docker running on your computer.

Additionally, make sure you have sourced the environment variables below. This can be done by putting the environment
variables in a file `.env`, and sourcing the environment variables using `source .env` in your terminal.

To start development, simply run `make` in the project directory. 

For more information, please read the [`Makefile`](./Makefile) or execute `make help` in the terminal.

#### Without `make`

To start a development server on your machine follow these steps:

1. Set local environment variables required to run the application:

   ```sh
   export POSTGRES_USER=admin
   export POSTGRES_PASSWORD=admin
   export POSTGRES_DB=a3s
   ```

2. Start your PostGreSQL server (`docker-compose up psql`)
3. Download Go dependencies with `go mod download`
4. Generate `gqlgen` code with `go run github.com/99designs/gqlgen generate`
5. Generate Ent entities with `go generate ./ent`
6. Migrate the Ent schema to your local PostGreSQL database with `go run cmd/migrate.go`
7. Start the server with `go run cmd/main.go`

[gqlgen]: https://github.com/99designs/gqlgen

[entgo]: https://entgo.io/
