package std_mysql

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMysqlDB(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", connectionString)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(2 * time.Minute)

	err = db.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
	if err != nil {
		return nil, err
	}

	return db, nil
}
