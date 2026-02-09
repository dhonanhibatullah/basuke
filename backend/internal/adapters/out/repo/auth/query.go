package adaptersoutrepoauth

import (
	"github.com/Masterminds/squirrel"
	domainmodel "github.com/dhonanhibatullah/basuke/internal/domain/model"
)

func (r *Repo) queryCreateAuth(userId int64, username string, email string, passwordHash string, by int64) (string, []any, error) {
	return r.sqr.
		Insert("basuke_auths").
		Columns(
			"user_id",
			"username",
			"password_hash",
			"email",
			"created_by",
		).
		Values(
			userId,
			username,
			passwordHash,
			email,
			by,
		).
		Suffix("RETURNING id").
		ToSql()
}

func (r *Repo) queryReadAuthByIdentifier(identifier string) (string, []any, error) {
	return r.sqr.
		Select(
			"id",
			"user_id",
			"username",
			"email",
			"password_hash",
			"status",
			"last_login_at",
			"created_at",
			"created_by",
			"updated_at",
			"updated_by",
		).
		From("basuke_auths").
		Where(squirrel.Or{
			squirrel.Eq{"username": identifier},
			squirrel.Eq{"email": identifier},
		}).
		ToSql()
}

func (r *Repo) queryReadAuthById(id int64) (string, []any, error) {
	return r.sqr.
		Select(
			"id",
			"user_id",
			"username",
			"email",
			"password_hash",
			"status",
			"last_login_at",
			"created_at",
			"created_by",
			"updated_at",
			"updated_by",
		).
		From("basuke_auths").
		Where(squirrel.Eq{"id": id}).
		ToSql()
}

func (r *Repo) queryUpdateAuth(id int64, username *string, email *string, by int64) (string, []any, error) {
	qb := r.sqr.
		Update("basuke_auths").
		Set("updated_at", squirrel.Expr("NOW()")).
		Set("updated_by", by).
		Where(squirrel.Eq{"id": id})

	changeCond := squirrel.Or{}

	if username != nil {
		qb = qb.Set("username", *username)
		changeCond = append(changeCond, squirrel.Or{
			squirrel.NotEq{"username": *username},
			squirrel.Eq{"username": nil},
		})
	}
	if email != nil {
		qb = qb.Set("email", *email)
		changeCond = append(changeCond, squirrel.Or{
			squirrel.NotEq{"email": *email},
			squirrel.Eq{"email": nil},
		})
	}

	if len(changeCond) > 0 {
		qb = qb.Where(changeCond)
	}

	return qb.ToSql()
}

func (r *Repo) queryUpdateAuthPasswordHash(id int64, passwordHash string, by int64) (string, []any, error) {
	return r.sqr.
		Update("basuke_auths").
		Set("password_hash", passwordHash).
		Set("updated_at", squirrel.Expr("NOW()")).
		Set("updated_by", by).
		Where(squirrel.Eq{"id": id}).
		ToSql()
}

func (r *Repo) queryUpdateAuthStatus(id int64, status domainmodel.AuthStatus, by int64) (string, []any, error) {
	return r.sqr.
		Update("basuke_auths").
		Set("status", status).
		Set("updated_at", squirrel.Expr("NOW()")).
		Set("updated_by", by).
		Where(squirrel.Eq{"id": id}).
		ToSql()
}
