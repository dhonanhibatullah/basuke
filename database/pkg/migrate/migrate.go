package migrate

import (
	"embed"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

type Direction uint8

const (
	DirectionUp Direction = iota
	DirectionDown
)

type Config struct {
	DBHost         string
	DBPort         uint16
	DBUsername     string
	DBPassword     string
	DBName         string
	Direction      Direction
	Step           uint
	MigrationFiles *embed.FS
}

// Logger implementation for migrate library
type migrateLogger struct{}

func (l *migrateLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l *migrateLogger) Verbose() bool {
	return true
}

func Migrate(config *Config) error {
	log.Println("Connecting to database...")
	db, err := connect(config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create database driver: %w", err)
	}
	sourceDriver, err := iofs.New(config.MigrationFiles, "sql")
	if err != nil {
		return fmt.Errorf("failed to create source driver: %w", err)
	}
	m, err := migrate.NewWithInstance("iofs", sourceDriver, "postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	// Set logger to see detailed output
	m.Log = &migrateLogger{}

	// Log pre-migration status
	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Printf("Failed to get current version: %v", err)
	} else {
		log.Printf("Current migration version: %v, Dirty: %v", version, dirty)
	}

	log.Printf("Starting migration (Direction: %v, Step: %d)...", config.Direction, config.Step)

	if config.Step > 0 {
		steps := int(config.Step)
		if config.Direction == DirectionDown {
			steps = -steps
		}
		if err := m.Steps(steps); err != nil {
			if err == migrate.ErrNoChange {
				log.Println("No migration changes needed.")
			} else {
				return fmt.Errorf("migration step failed: %w", err)
			}
		} else {
			log.Println("Migration step completed successfully.")
		}
	} else {
		if config.Direction == DirectionUp {
			if err := m.Up(); err != nil {
				if err == migrate.ErrNoChange {
					log.Println("No migration changes needed.")
				} else {
					return fmt.Errorf("migration up failed: %w", err)
				}
			} else {
				log.Println("Migration up completed successfully.")
			}
		} else {
			if err := m.Down(); err != nil {
				if err == migrate.ErrNoChange {
					log.Println("No migration changes needed.")
				} else {
					return fmt.Errorf("migration down failed: %w", err)
				}
			} else {
				log.Println("Migration down completed successfully.")
			}
		}
	}

	// Log post-migration status
	version, dirty, err = m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		log.Printf("Failed to get new version: %v", err)
	} else {
		log.Printf("New migration version: %v, Dirty: %v", version, dirty)
	}

	return nil
}

func (d Direction) String() string {
	if d == DirectionUp {
		return "Up"
	}
	return "Down"
}
