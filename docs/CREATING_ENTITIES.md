# Creating Entities

To create a new Entity in the project, you need to do two things:

1. Create a user-facing GraphQL type
2. Create an internal Ent schema entity

## Creating a GraphQL type

To keep schemas as simple as possible, we create a `.graphqls` file for each
top-level entity. See the existing schemas in [`/graph`](/graph) for examples.

## Creating an Ent entity

Ent is able to generate `.go` files for the schema and set up everything, except
for fields and edges. To create a new entity, run the following command:

```sh
# Create a new entity
make entity name=<EntityName>
```

The generated file is found in [`/ent/schema`](/ent/schema).
