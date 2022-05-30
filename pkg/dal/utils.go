package dal

import (
	"database/sql"
	"highload-architect/pkg/config"
)

func getDbConnect() (*sql.DB, error) {
	db, dbErr := sql.Open("mysql", config.GetDbConfig())
	if dbErr != nil {
		println("getDbConnect ", dbErr)
	}

	return db, dbErr
}
