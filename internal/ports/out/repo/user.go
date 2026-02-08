package portsoutrepo

import (
	"context"
	"encoding/json"
	"time"

	domainmodel "github.com/dhonanhibatullah/basuke/internal/domain/model"
)

type User interface {
	CreateUser(ctx context.Context, name string, role domainmodel.UserRole, by int64) (id int64, err error)
	CreateUserSelf(ctx context.Context, name string, role domainmodel.UserRole) (id int64, err error)
	ReadUsersByPagination(ctx context.Context, search *string, page *int, limit *int, cursorId *int64, role *domainmodel.UserRole, status *domainmodel.UserStatus) (users []domainmodel.User, total int64, err error)
	ReadUserById(ctx context.Context, id int64) (user *domainmodel.User, err error)
	ReadUserByAuthId(ctx context.Context, authId int64) (user *domainmodel.User, err error)
	UpdateUser(ctx context.Context, id int64, name *string, role *domainmodel.UserRole, bio *string, preferences *json.RawMessage, by int64) (err error)
	UpdateUserStatus(ctx context.Context, id int64, status domainmodel.UserStatus, by int64) (err error)
	UpdateUserLastLoginAt(ctx context.Context, id int64) (err error)
	UpdateUsersRequiredToOffline(ctx context.Context, cutoffDuration time.Duration) (total int64, err error)
	DeleteUserById(ctx context.Context, id int64) (err error)
}
