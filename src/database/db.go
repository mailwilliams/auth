package database

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewSQLConnection(ctx context.Context) (*sql.DB, error) {
	var (
		db  *sql.DB
		err error
	)

	//	opening new connection to database
	//	TODO: @Matt change dataSourceName to be environment variables using docker

	db, err = sql.Open("mysql", "root:password@tcp(db:3306)/auth")
	if err != nil {
		return db, err
	}

	return db, err
}
