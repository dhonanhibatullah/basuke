package domainmodel

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccountRole string

const (
	AccountRoleSuper  AccountRole = "super"
	AccountRoleAdmin  AccountRole = "admin"
	AccountRoleClient AccountRole = "client"
)

type AccountStatus string

const (
	AccountStatusActive    AccountStatus = "active"
	AccountStatusInactive  AccountStatus = "inactive"
	AccountStatusSuspended AccountStatus = "suspended"
)

type Account struct {
	Id            int64
	Username      string
	Email         string
	EmailVerified bool
	PasswordHash  string
	Name          string
	Role          AccountRole
	Status        AccountStatus
	Bio           *string
	Preferences   *json.RawMessage
	LastLoginAt   *time.Time
	CreatedAt     *time.Time
	CreatedBy     *int64
	UpdatedAt     *time.Time
	UpdatedBy     *int64
}

type AccountGenericClaims struct {
	AccountId int64       `json:"account_id"`
	Username  string      `json:"username"`
	Name      string      `json:"name"`
	Role      AccountRole `json:"role"`
}

type AccountJwtClaims struct {
	AccountGenericClaims
	jwt.RegisteredClaims
}
