package domain

import "github.com/google/uuid"

type Client struct {
	Id             uuid.UUID
	Name           string
	EmailAddress   string
	CategoryId     uuid.UUID
	WorkScheduleId uuid.UUID
}

type ClientResponse struct {
	Client       Client
	Category     Category
	WorkSchedule WorkSchedule
	Specialties  []Specialty
}
