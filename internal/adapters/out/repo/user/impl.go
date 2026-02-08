package adaptersoutrepouser

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/Masterminds/squirrel"
	domainmodel "github.com/dhonanhibatullah/basuke/internal/domain/model"
	portsoutrepo "github.com/dhonanhibatullah/basuke/internal/ports/out/repo"
	"github.com/dhonanhibatullah/basuke/pkg/pgxdx"
	"github.com/jackc/pgx/v5"
)

type Repo struct {
	dx  pgxdx.Pgxdx
	sqr *squirrel.StatementBuilderType
}

func (r *Repo) New(dx pgxdx.Pgxdx, sqr *squirrel.StatementBuilderType) portsoutrepo.User {
	return &Repo{
		dx:  dx,
		sqr: sqr,
	}
}

func (r *Repo) CreateUser(ctx context.Context, name string, role domainmodel.UserRole, by int64) (id int64, err error) {
	query, args, err := r.queryCreateUser(name, role, by)
	if err != nil {
		return 0, err
	}

	err = r.dx.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, r.scanErr(err)
	}

	return
}

func (r *Repo) CreateUserSelf(ctx context.Context, name string, role domainmodel.UserRole) (id int64, err error) {
	err = r.dx.QueryRow(ctx, r.queryGenerateNewId()).Scan(&id)
	if err != nil {
		return 0, r.scanErr(err)
	}

	query, args, err := r.queryCreateUserSelf(id, name, role)
	if err != nil {
		return 0, err
	}

	err = r.dx.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, r.scanErr(err)
	}

	return
}

func (r *Repo) ReadUsersByPagination(ctx context.Context, search *string, page *int, limit *int, cursorId *int64, role *domainmodel.UserRole, status *domainmodel.UserStatus) (users []domainmodel.User, total int64, err error) {
	countQuery, countArgs, query, args, err := r.queryReadUserByPagination(search, page, limit, cursorId, role, status)
	if err != nil {
		return nil, 0, err
	}

	var totalItem int
	err = r.dx.QueryRow(ctx, countQuery, countArgs...).Scan(&totalItem)
	if err != nil {
		return nil, 0, r.scanErr(err)
	}
	total = int64(totalItem)

	rows, err := r.dx.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, r.scanErr(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user domainmodel.User
		if err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Role,
			&user.Status,
		); err != nil {
			return nil, 0, r.scanErr(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, r.scanErr(err)
	}

	return
}

func (r *Repo) ReadUserById(ctx context.Context, id int64) (user *domainmodel.User, err error) {
	query, args, err := r.queryReadUserById(id)
	if err != nil {
		return nil, err
	}

	user = &domainmodel.User{}
	err = r.dx.QueryRow(ctx, query, args...).Scan(
		&user.Id,
		&user.Name,
		&user.Role,
		&user.Status,
		&user.Bio,
		&user.LastLoginAt,
		&user.Preferences,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.UpdatedAt,
		&user.UpdatedBy,
	)
	if err != nil {
		return nil, r.scanErr(err)
	}

	return
}

func (r *Repo) ReadUserByAuthId(ctx context.Context, authId int64) (user *domainmodel.User, err error) {
	query, args, err := r.queryReadUserByAuthId(authId)
	if err != nil {
		return nil, err
	}

	user = &domainmodel.User{}
	err = r.dx.QueryRow(ctx, query, args...).Scan(
		&user.Id,
		&user.Name,
		&user.Role,
		&user.Status,
		&user.Bio,
		&user.LastLoginAt,
		&user.Preferences,
		&user.CreatedAt,
		&user.CreatedBy,
		&user.UpdatedAt,
		&user.UpdatedBy,
	)
	if err != nil {
		return nil, r.scanErr(err)
	}

	return
}

func (r *Repo) UpdateUser(ctx context.Context, id int64, name *string, role *domainmodel.UserRole, bio *string, preferences *json.RawMessage, by int64) (err error) {
	query, args, err := r.queryUpdateUser(id, name, role, bio, preferences, by)
	if err != nil {
		return err
	}

	cmd, err := r.dx.Exec(ctx, query, args...)
	if err != nil {
		return r.scanErr(err)
	}
	if cmd.RowsAffected() == 0 {
		return domainmodel.ErrNoChange
	}

	return
}

func (r *Repo) UpdateUserStatus(ctx context.Context, id int64, status domainmodel.UserStatus, by int64) (err error) {
	query, args, err := r.queryUpdateUserStatus(id, status, by)
	if err != nil {
		return err
	}

	cmd, err := r.dx.Exec(ctx, query, args...)
	if err != nil {
		return r.scanErr(err)
	}
	if cmd.RowsAffected() == 0 {
		return domainmodel.ErrNoChange
	}

	return
}

func (r *Repo) UpdateUserLastLoginAt(ctx context.Context, id int64) (err error) {
	query, args, err := r.queryUpdateUserLastLoginAt(id)
	if err != nil {
		return err
	}

	cmd, err := r.dx.Exec(ctx, query, args...)
	if err != nil {
		return r.scanErr(err)
	}
	if cmd.RowsAffected() == 0 {
		return domainmodel.ErrNoChange
	}

	return
}

func (r *Repo) UpdateUsersRequiredToOffline(ctx context.Context, cutoffDuration time.Duration) (total int64, err error) {
	cutoffTime := time.Now().Add(-cutoffDuration)
	query, args, err := r.queryUpdateUsersRequiredToOffline(cutoffTime)
	if err != nil {
		return 0, err
	}

	cmd, err := r.dx.Exec(ctx, query, args...)
	if err != nil {
		return 0, r.scanErr(err)
	}

	total = cmd.RowsAffected()

	return
}

func (r *Repo) DeleteUserById(ctx context.Context, id int64) (err error) {
	query, args, err := r.queryDeleteUser(id)
	if err != nil {
		return err
	}

	cmd, err := r.dx.Exec(ctx, query, args...)
	if err != nil {
		return r.scanErr(err)
	}
	if cmd.RowsAffected() == 0 {
		return domainmodel.ErrNoChange
	}

	return
}

func (r *Repo) scanErr(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return domainmodel.ErrNotFound
	}

	return err
}
