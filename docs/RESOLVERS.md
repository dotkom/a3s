# Resolvers documentation

## Implementation

See the [`resolvers`](../resolvers) directory for the implementation of the
GraphQL resolvers.

## Conventions

GraphQL resolvers are automatically generated from
[`graph/schema.graphqls`](../graph/schema.graphqls) in
[`resolvers/schema.resolvers.go`](../resolvers/schema.resolvers.go).

In order to maintain a clean
[`resolvers/schema.resolvers.go`](../resolvers/schema.resolvers.go) file, the
implementation logic for each resolver is split into each own file corresponding
to a GraphQL type. The file is named after the GraphQL type, with the
`.resolvers.go` file extension.

For instance, for an `Event`, there exists a resolver
[`resolvers/event.resolvers.go`](../resolvers/event.resolvers.go) with the
implementation detail concerning an event. In
[`resolvers/schema.resolvers.go`](../resolvers/schema.resolvers.go), we simply
use the relevant function from
[`resolvers/event.resolvers.go`](../resolvers/event.resolvers.go).

**In short, take a look at some existing resolver and try to notice the pattern
of naming and how to resolvers are used in
[`resolvers/schema.resolvers.go`](../resolvers/schema.resolvers.go)**.
