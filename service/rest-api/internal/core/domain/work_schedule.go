package domain

import "github.com/google/uuid"

type WorkSchedule struct {
	Id        uuid.UUID
	StartTime string
	EndTime   string
}
