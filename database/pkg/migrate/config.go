package migrate

import (
	"embed"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func envGetString(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("environment variable %s is required", key)
	}
	return val, nil
}

func envGetUint(key string) (uint, error) {
	val, err := envGetString(key)
	if err != nil {
		return 0, err
	}
	valUint, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(valUint), nil
}

func envGetUint16(key string) (uint16, error) {
	val, err := envGetUint(key)
	if err != nil {
		return 0, err
	} else if val > 65535 {
		return 0, fmt.Errorf("environment variable %s must be less than or equal to 65535", key)
	}
	return uint16(val), nil
}

func envGetDirection(key string) (Direction, error) {
	val, err := envGetString(key)
	if err != nil {
		return DirectionUp, err
	}
	switch val {
	case "up":
		return DirectionUp, nil
	case "down":
		return DirectionDown, nil
	default:
		return DirectionUp, fmt.Errorf("environment variable %s must be either 'up' or 'down'", key)
	}
}

func LoadConfig(path string, files *embed.FS) (*Config, error) {
	config := &Config{}

	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	config.DBHost, err = envGetString("MIGRATE_DB_HOST")
	if err != nil {
		return nil, err
	}
	config.DBPort, err = envGetUint16("MIGRATE_DB_PORT")
	if err != nil {
		return nil, err
	}
	config.DBUsername, err = envGetString("MIGRATE_DB_USERNAME")
	if err != nil {
		return nil, err
	}
	config.DBPassword, err = envGetString("MIGRATE_DB_PASSWORD")
	if err != nil {
		return nil, err
	}
	config.DBName, err = envGetString("MIGRATE_DB_NAME")
	if err != nil {
		return nil, err
	}
	config.Direction, err = envGetDirection("MIGRATE_DB_DIRECTION")
	if err != nil {
		return nil, err
	}
	config.Step, err = envGetUint("MIGRATE_DB_STEP")
	if err != nil {
		return nil, err
	}
	config.MigrationFiles = files

	return config, nil
}
