package postgres

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/rs/zerolog"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DBConnect struct {
	DBConnect *sql.DB
	Logger    *zerolog.Logger
}

// GetNewConnection создаёт новое соединение с базой данных.
func GetNewConnection(db *sql.DB, dbConf string, logger *zerolog.Logger) *DBConnect {

	migration, err := migrate.New("file://migrations", dbConf)
	if err != nil {
		logger.Error().Msg(err.Error())
	}

	if err = migration.Up(); err != nil {
		logger.Error().Msg(err.Error())
	}

	return &DBConnect{DBConnect: db, Logger: logger}
}
