package resolvers

import "github.com/dotkom/a3s/repository"

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	EventOrganizerRepository *repository.EventOrganizerRepository
}
