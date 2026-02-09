package adaptersoutrepouser

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	domainmodel "github.com/dhonanhibatullah/basuke/internal/domain/model"
)

func (r *Repo) queryCreateUser(name string, role domainmodel.UserRole, by int64) (string, []any, error) {
	return r.sqr.
		Insert("basuke_users").
		Columns(
			"name",
			"role",
			"created_by",
		).
		Values(
			name,
			role,
			by,
		).
		Suffix("RETURNING id").
		ToSql()
}

func (r *Repo) queryCreateUserSelf(id int64, name string, role domainmodel.UserRole) (string, []any, error) {
	return r.sqr.
		Insert("basuke_users").
		Columns(
			"id",
			"name",
			"role",
			"created_by",
		).
		Values(
			id,
			name,
			role,
			id,
		).
		Suffix("RETURNING id").
		ToSql()
}

func (r *Repo) queryReadUserByPagination(search *string, page *int, limit *int, cursorId *int64, role *domainmodel.UserRole, status *domainmodel.UserStatus) (string, []any, string, []any, error) {
	countQb := r.sqr.Select("COUNT(id)").From("basuke_users")
	qb := r.sqr.
		Select(
			"id",
			"name",
			"role",
			"status",
		).
		From("basuke_users")

	if search != nil {
		likePattern := fmt.Sprintf("%%%s%%", *search)
		clause := squirrel.Or{squirrel.ILike{"name": likePattern}}
		countQb = countQb.Where(clause)
		qb = qb.Where(clause)
	}
	if role != nil {
		countQb = countQb.Where(squirrel.Eq{"role": *role})
		qb = qb.Where(squirrel.Eq{"role": *role})
	}
	if status != nil {
		countQb = countQb.Where(squirrel.Eq{"status": *status})
		qb = qb.Where(squirrel.Eq{"status": *status})
	}
	if cursorId != nil {
		countQb = countQb.Where(squirrel.Lt{"id": *cursorId})
		qb = qb.Where(squirrel.Lt{"id": *cursorId})
	}

	vPage := 1
	vLimit := 10
	if page != nil {
		vPage = *page
	}
	if limit != nil {
		vLimit = *limit
	}
	qb = qb.OrderBy("id DESC").Limit(uint64(vLimit)).Offset(uint64((vPage - 1) * vLimit))

	countQuery, countArgs, err := countQb.ToSql()
	if err != nil {
		return "", nil, "", nil, err
	}
	query, args, err := qb.ToSql()
	if err != nil {
		return "", nil, "", nil, err
	}

	return countQuery, countArgs, query, args, nil
}

func (r *Repo) queryReadUserById(id int64) (string, []any, error) {
	return r.sqr.
		Select(
			"id",
			"name",
			"role",
			"status",
			"bio",
			"preferences",
			"created_at",
			"created_by",
			"updated_at",
			"updated_by",
		).
		From("basuke_users").
		Where(squirrel.Eq{"id": id}).
		ToSql()
}

func (r *Repo) queryReadUserByAuthId(authId int64) (string, []any, error) {
	return r.sqr.
		Select(
			"u.id",
			"u.name",
			"u.role",
			"u.status",
			"u.bio",
			"u.preferences",
			"u.created_at",
			"u.created_by",
			"u.updated_at",
			"u.updated_by",
		).
		From("basuke_users u").
		Join("basuke_auths a ON a.user_id = u.id").
		Where(squirrel.Eq{"a.id": authId}).
		ToSql()
}

func (r *Repo) queryUpdateUser(id int64, name *string, bio *string, preferences *json.RawMessage, by int64) (string, []any, error) {
	qb := r.sqr.
		Update("basuke_users").
		Set("updated_by", by).
		Set("updated_at", squirrel.Expr("NOW()")).
		Where(squirrel.Eq{"id": id})

	changeCond := squirrel.Or{}

	if name != nil {
		qb = qb.Set("name", *name)
		changeCond = append(changeCond, squirrel.Or{
			squirrel.NotEq{"name": *name},
			squirrel.Eq{"name": nil},
		})
	}
	if bio != nil {
		qb = qb.Set("bio", *bio)
		changeCond = append(changeCond, squirrel.Or{
			squirrel.NotEq{"bio": *bio},
			squirrel.Eq{"bio": nil},
		})
	}
	if preferences != nil {
		qb = qb.Set("preferences", preferences)
		changeCond = append(changeCond, squirrel.Or{
			squirrel.NotEq{"preferences": preferences},
			squirrel.Eq{"preferences": nil},
		})
	}

	if len(changeCond) > 0 {
		qb = qb.Where(changeCond)
	}

	return qb.ToSql()
}

func (r *Repo) queryUpdateUserRole(id int64, role domainmodel.UserRole, by int64) (string, []any, error) {
	return r.sqr.
		Update("basuke_users").
		Set("role", role).
		Set("updated_at", squirrel.Expr("NOW()")).
		Set("updated_by", by).
		Where(squirrel.Eq{"id": id}).
		ToSql()
}

func (r *Repo) queryUpdateUserStatus(id int64, status domainmodel.UserStatus, by int64) (string, []any, error) {
	return r.sqr.
		Update("basuke_users").
		Set("status", status).
		Set("updated_at", squirrel.Expr("NOW()")).
		Set("updated_by", by).
		Where(squirrel.Eq{"id": id}).
		ToSql()
}

func (r *Repo) queryUpdateUserLastLoginAt(id int64) (string, []any, error) {
	return r.sqr.
		Update("basuke_users").
		Set("last_login_at", squirrel.Expr("NOW()")).
		Set("updated_at", squirrel.Expr("NOW()")).
		Set("updated_by", id).
		Where(squirrel.Eq{"id": id}).
		ToSql()
}

func (r *Repo) queryUpdateUsersRequiredToOffline(cutoffTime time.Time) (string, []any, error) {
	subQuery, subArgs, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question).
		Select("user_id").
		From("basuke_auths").
		Where(squirrel.Lt{"last_login_at": cutoffTime}).
		ToSql()
	if err != nil {
		return "", nil, err
	}

	return r.sqr.
		Update("basuke_users").
		Set("status", domainmodel.UserStatusOffline).
		Set("updated_at", squirrel.Expr("NOW()")).
		Set("updated_by", squirrel.Expr("id")).
		Where(squirrel.Eq{"status": domainmodel.UserStatusOnline}).
		Where("id IN ("+subQuery+")", subArgs...).
		ToSql()
}

func (r *Repo) queryDeleteUser(id int64) (string, []any, error) {
	return r.sqr.
		Delete("basuke_users").
		Where(squirrel.Eq{"id": id}).
		ToSql()
}

func (r *Repo) queryGenerateNewId() string {
	return "SELECT nextval(pg_get_serial_sequence('basuke_users', 'id'));"
}
