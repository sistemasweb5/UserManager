package domain

import "github.com/google/uuid"

type Specialty struct {
	Id     uuid.UUID
	Name   string
	ClientId uuid.UUID
}
