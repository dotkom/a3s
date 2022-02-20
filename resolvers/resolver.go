package resolvers

import "github.com/dotkom/a3s/ent"

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	Client *ent.Client
}
