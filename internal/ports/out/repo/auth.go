package portsoutrepo

import (
	"context"

	domainmodel "github.com/dhonanhibatullah/basuke/internal/domain/model"
)

type Auth interface {
	CreateAuth(ctx context.Context, userId int64, username string, passwordHash string, email *string, by int64) (id int64, err error)
	ReadAuthByIdentifier(ctx context.Context, identifier string) (auth *domainmodel.Auth, err error)
	ReadAuthById(ctx context.Context, id int64) (auth *domainmodel.Auth, err error)
	UpdateAuth(ctx context.Context, id int64, username *string, email *string, by int64) (err error)
	UpdateAuthPasswordHash(ctx context.Context, id int64, passwordHash string, by int64) (err error)
	UpdateAuthStatus(ctx context.Context, id int64, status domainmodel.AuthStatus, by int64) (err error)
}
