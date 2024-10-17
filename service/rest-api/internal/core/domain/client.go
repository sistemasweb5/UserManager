package domain

import "github.com/google/uuid"

type Client struct {
	Name             string
	MembershipNumber uuid.UUID
	EmailAddress     string
}
