package adaptersoutrepoauth

import (
	"context"
	"errors"
	"strings"

	"github.com/Masterminds/squirrel"
	domainmodel "github.com/dhonanhibatullah/basuke/internal/domain/model"
	portsoutrepo "github.com/dhonanhibatullah/basuke/internal/ports/out/repo"
	"github.com/dhonanhibatullah/basuke/pkg/pgxdx"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Repo struct {
	dx  pgxdx.Pgxdx
	sqr *squirrel.StatementBuilderType
}

func (r *Repo) New(dx pgxdx.Pgxdx, sqr *squirrel.StatementBuilderType) portsoutrepo.Auth {
	return &Repo{
		dx:  dx,
		sqr: sqr,
	}
}

func (r *Repo) CreateAuth(ctx context.Context, userId int64, username string, email string, passwordHash string, by int64) (id int64, err error) {
	query, args, err := r.queryCreateAuth(userId, username, email, passwordHash, by)
	if err != nil {
		return 0, err
	}

	err = r.dx.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, r.scanErr(err)
	}

	return
}

func (r *Repo) ReadAuthByIdentifier(ctx context.Context, identifier string) (auth *domainmodel.Auth, err error) {
	query, args, err := r.queryReadAuthByIdentifier(identifier)
	if err != nil {
		return nil, err
	}

	auth = &domainmodel.Auth{}
	err = r.dx.QueryRow(ctx, query, args...).Scan(
		&auth.Id,
		&auth.UserId,
		&auth.Username,
		&auth.Email,
		&auth.PasswordHash,
		&auth.Status,
		&auth.LastLoginAt,
		&auth.CreatedAt,
		&auth.CreatedBy,
		&auth.UpdatedAt,
		&auth.UpdatedBy,
	)
	if err != nil {
		return nil, r.scanErr(err)
	}

	return
}

func (r *Repo) ReadAuthById(ctx context.Context, id int64) (auth *domainmodel.Auth, err error) {
	query, args, err := r.queryReadAuthById(id)
	if err != nil {
		return nil, err
	}

	auth = &domainmodel.Auth{}
	err = r.dx.QueryRow(ctx, query, args...).Scan(
		&auth.Id,
		&auth.UserId,
		&auth.Username,
		&auth.Email,
		&auth.PasswordHash,
		&auth.Status,
		&auth.LastLoginAt,
		&auth.CreatedAt,
		&auth.CreatedBy,
		&auth.UpdatedAt,
		&auth.UpdatedBy,
	)
	if err != nil {
		return nil, r.scanErr(err)
	}

	return
}

func (r *Repo) UpdateAuth(ctx context.Context, id int64, username *string, email *string, by int64) (err error) {
	query, args, err := r.queryUpdateAuth(id, username, email, by)
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

func (r *Repo) UpdateAuthPasswordHash(ctx context.Context, id int64, passwordHash string, by int64) (err error) {
	query, args, err := r.queryUpdateAuthPasswordHash(id, passwordHash, by)
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

func (r *Repo) UpdateAuthStatus(ctx context.Context, id int64, status domainmodel.AuthStatus, by int64) (err error) {
	query, args, err := r.queryUpdateAuthStatus(id, status, by)
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

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			if strings.Contains(pgErr.ConstraintName, "username") {
				return domainmodel.ErrUsernameExist
			}
			if strings.Contains(pgErr.ConstraintName, "email") {
				return domainmodel.ErrEmailExist
			}
		}
	}

	return err
}
