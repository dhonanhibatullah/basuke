package database

import (
	"github.com/dhonanhibatullah/basuke/database/migrations"
	"github.com/dhonanhibatullah/basuke/database/pkg/migrate"
)

func main() {
	cfg, err := migrate.LoadConfig(".env", &migrations.Sqls)
	if err != nil {
		panic(err)
	}
	err = migrate.Migrate(cfg)
	if err != nil {
		panic(err)
	}
}
