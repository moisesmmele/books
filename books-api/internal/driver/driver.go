package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"time"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenConns = 10
const maxIdleConns = 5
const maxConnLifetime = 5 * time.Minute

func ConnectPostgres(dsn string) (*DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(maxConnLifetime)

	err = testDB(db)

	dbConn.SQL = db
	return dbConn, err
}

func testDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to database", err)
		return err
	}
	fmt.Println("Successfully connected to database")
	return nil
}
