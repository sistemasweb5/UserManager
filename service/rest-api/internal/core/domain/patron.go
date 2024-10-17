package domain

import "github.com/google/uuid"

type User struct {
	Name             string
	MembershipNumber uuid.UUID
	EmailAddress     string
}
