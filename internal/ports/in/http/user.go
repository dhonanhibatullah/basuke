package portsinhttp

import (
	"context"

	domainmodel "github.com/dhonanhibatullah/basuke/internal/domain/model"
)

type User interface {
	GetUsersList(ctx context.Context, search *string, page *int, limit *int, cursorId *int64, role *domainmodel.UserRole, status *domainmodel.UserStatus) (users []User, err error)
	GetUserById(ctx context.Context, id string) (user *domainmodel.User, err error)
	UpdateUser(ctx context.Context, id string, name string, username string, email string, password string) (err error)
	DeleteUserById(ctx context.Context, id string) (err error)
}
