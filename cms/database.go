package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/MikelSot/metal-bat/model"
)

func newDBConnection(dbConf model.Database) (*sql.DB, error) {
	if dbConf.SSLMode == "" {
		dbConf.SSLMode = "disable"
	}

	dns := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		dbConf.User,
		dbConf.Password,
		dbConf.Server,
		dbConf.Port,
		dbConf.Name,
		dbConf.SSLMode,
	)

	return sql.Open("postgres", dns)
}
