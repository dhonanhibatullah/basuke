package domainmodel

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserRole string

const (
	UserRoleSuper  UserRole = "super"
	UserRoleAdmin  UserRole = "admin"
	UserRoleClient UserRole = "client"
)

type UserStatus string

const (
	UserStatusOnline  UserStatus = "online"
	UserStatusOffline UserStatus = "offline"
	UserStatusAway    UserStatus = "away"
	UserStatusDnd     UserStatus = "dnd"
)

type User struct {
	Id          int64
	Name        string
	Role        UserRole
	Status      UserStatus
	Bio         *string
	LastLoginAt *time.Time
	Preferences *json.RawMessage
	CreatedAt   *time.Time
	CreatedBy   *int64
	UpdatedAt   *time.Time
	UpdatedBy   *int64
}

type UserGenericClaims struct {
	UserId   int64    `json:"user_id"`
	AuthId   int64    `json:"auth_id"`
	Username string   `json:"username"`
	Name     string   `json:"name"`
	Role     UserRole `json:"role"`
}

type UserApiKeyClaims struct {
	UserGenericClaims
	CreatedAt time.Time `json:"created_at"`
}

type UserJwtClaims struct {
	UserGenericClaims
	jwt.RegisteredClaims
}
