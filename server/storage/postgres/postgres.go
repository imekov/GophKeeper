package postgres

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DBConnect struct {
	DBConnect *sql.DB
}

func GetNewConnection(db *sql.DB, dbConf string) *DBConnect {

	migration, err := migrate.New("file://migrations", dbConf)
	if err != nil {
		log.Print(err)
	}

	if err = migration.Up(); err != nil {
		log.Print(err)
	}

	return &DBConnect{DBConnect: db}
}
