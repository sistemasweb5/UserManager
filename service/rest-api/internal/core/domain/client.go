package domain

import "github.com/google/uuid"

type Worker struct {
	Id             uuid.UUID
	Name           string
	EmailAddress   string
	WorkScheduleId uuid.UUID
}

type WorkerResponse struct {
	Id           uuid.UUID
	Name         string
	EmailAddress string
	WorkSchedule WorkSchedule
	Specialties  []Specialty
}

type Applicant struct {
	Id           uuid.UUID
	Name         string
	EmailAddress string
}

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

type ApplicantRequest struct {
	Name         string
	EmailAddress string
}

type WorkerRequest struct {
	Name           string
	EmailAddress   string
	WorkScheduleId uuid.UUID
}
