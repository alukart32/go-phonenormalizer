//go:build migrate

package app

import (
	"errors"
	"fmt"
	"log"
	"time"

	"alukart32.com/phoneNormalizer/config"
	"github.com/golang-migrate/migrate/v4"

	// migrate tools
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	_defaultAttempts = 5
	_defaultTimeout  = time.Second
	_configPath      = "../../config/config.yaml"
	_migrationsPath  = "../../migrations"
)

func init() {
	cfg, _ := config.NewConfig(_configPath)
	// dbdriver://username:password@host:port/dbname?param1=true&param2=false
	dbURL := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DB.Driver, cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DBname)

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://"+_migrationsPath, dbURL)
		if err == nil {
			break
		}

		log.Printf("migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("migrate: postgres connect error: %s", err)
	}

	err = m.Up()
	defer m.Close()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("migrate: no change...")
		return
	}

	log.Printf("migrate db: up success...")
}
