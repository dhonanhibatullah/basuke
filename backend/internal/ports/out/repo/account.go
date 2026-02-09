package portsoutrepo

import (
	"context"
	"encoding/json"

	domainmodel "github.com/dhonanhibatullah/basuke/backend/internal/domain/model"
)

type Account interface {
	CreateAccount(
		ctx context.Context,
		name string,
		username string,
		email string,
		passwordHash string,
		role domainmodel.AccountRole,
	) (id int64, err error)

	ReadAccountsByPagination(
		ctx context.Context,
		page int64,
		limit int64,
		search *string,
		role *domainmodel.AccountRole,
		status *domainmodel.AccountStatus,
	) (accounts []domainmodel.Account, totalData int64, err error)

	ReadAccountById(
		ctx context.Context,
		id int64,
	) (account *domainmodel.Account, err error)

	ReadAccountByIdentifier(
		ctx context.Context,
		identifier string,
	) (account *domainmodel.Account, err error)

	UpdateAccount(
		ctx context.Context,
		id int64,
		name *string,
		email *string,
		bio *string,
	) (err error)

	UpdateAccountPassword(
		ctx context.Context,
		id int64,
		passwordHash string,
	) (err error)

	UpdateAccountRole(
		ctx context.Context,
		id int64,
		role domainmodel.AccountRole,
	) (err error)

	UpdateAccountStatus(
		ctx context.Context,
		id int64,
		status domainmodel.AccountStatus,
	) (err error)

	UpdateAccountPreferences(
		ctx context.Context,
		id int64,
		preferences *json.RawMessage,
	) (err error)

	DeleteAccountById(
		ctx context.Context,
		id int64,
	) (err error)
}
