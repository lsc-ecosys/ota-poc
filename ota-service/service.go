package ota

import (
	"context"
	"errors"
)

var (
	// ErrUnautorizedAccess indicates missing or invalid credentials provided
	// when accessing a protected resource.
	ErrUnauthorizedAccess = errors.New("missing or invalid credentials provided")

	// ErrNotFound indicates a non-existent entity request.
	ErrNotFound = errors.New("non-existent entity")
)

// Service specifies an API that must be fullfiled by the domain service
// implementation, and all of its decorators (e.g. logging & metrics).
type Service interface {
	// for management APIs
	// CreateRollout adds a deployment action including artifact version and distribution criterias
	CreateRollout(context.Context, string, Rollout) (Rollout, error)

	// For device APIs
	// CheckUpToDate should check incoming request of device if matches the update criteria then return rollout or error
	CheckUpToDate(context.Context, string) (Rollout, error)

	// CreateNotification()
}
