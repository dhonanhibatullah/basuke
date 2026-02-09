package migrate

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func connect(host string, port uint16, user string, pass string, name string) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, pass, host, port, name)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
