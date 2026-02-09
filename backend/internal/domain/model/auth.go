package domainmodel

import "time"

type AuthStatus string

const (
	AuthStatusActivated   AuthStatus = "activated"
	AuthStatusDeactivated AuthStatus = "deactivated"
	AuthStatusSuspended   AuthStatus = "suspended"
)

type Auth struct {
	Id            int64
	UserId        int64
	Username      string
	Email         string
	EmailVerified bool
	PasswordHash  string
	Status        AuthStatus
	LastLoginAt   *time.Time
	CreatedAt     *time.Time
	CreatedBy     *int64
	UpdatedAt     *time.Time
	UpdatedBy     *int64
}
