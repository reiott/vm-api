package vmapi

import (
	"context"
	"time"
)

const (
	ErrVMNotFound = Error("vm not found")
)

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	StatusCreating = Status("Creating")
	StatusActive   = Status("Active")
	StatusDeleting = Status("Deleting")
)

type Status string

type VM struct {
	ID               string    `json:"id"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	Status           Status    `json:"status"`
	AvailabilityZone string    `json:"availabilityZone"`
	Hostname         string    `json:"hostname"`
	DevAdmin         string    `json:"devAdmin"`
	Flavor           string    `json:"flavor"`
	Image            string    `json:"image"`
}

type VMStore interface {
	All(context.Context) ([]*VM, error)
	Get(context.Context, string) (*VM, error)
	Update(context.Context, *VM) error
	Delete(context.Context, string) error
	SetStatus(context.Context, string, Status) error
}
