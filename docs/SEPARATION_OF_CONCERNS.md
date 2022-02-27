# Separation of concerns

The logic in this service is separated into the following layers:

- GraphQL layer
- Server layer
- Database layer

## GraphQL layer

This layer is responsible for the logic related to the GraphQL API, the `.go`
resolvers that are generated based on the GraphQL schema, and its implementation
logic.

See [`/graph`](/graph) for the GraphQL schemas.

See [`/resolvers/schema.resolvers.go`](/resolvers/schema.resolvers.go) for the
main resolver. This resolver delegates the responsibility of implementing logic
related to GraphQL to other specific resolvers located in the
[`/resolvers`](/resolvers) directory.

## Server layer

This layer is responsible for the logic related to the authentication,
validation, and other responsibilities that the GraphQL layer should not know
about.

See the [`/resolvers`](/resolvers) directory for the specific resolvers.

## Database layer

This layer is responsible for interacting with the database.

See the [`/repository`](/repository) directory for the repositories handling
interaction with the database.

## In conclusion

This separation of concerns makes it easier to work on a particular part of the
code base, in addition to facilitating testing units of code in isolation in an
simple manner.
