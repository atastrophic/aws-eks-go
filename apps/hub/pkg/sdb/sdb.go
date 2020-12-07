package sdb

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/atastrophic/go-ms-with-eks/pkg/appconfig"
	"github.com/atastrophic/go-ms-with-eks/pkg/exception"
	"github.com/jmoiron/sqlx"
)

type SqlDb struct {
	sql  *sql.DB
	sqlx *sqlx.DB
}

func NewSDB(config appconfig.SQL) *SqlDb {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", config.DSN.Username, config.DSN.Password, config.DSN.Protocol, config.DSN.URL, config.DSN.Port, config.DSN.Database)
	db, err := sql.Open("mysql", dsn)
	exception.WithError(err)

	err = db.Ping()
	exception.WithError(err)

	sqlxDb := sqlx.NewDb(db, "mysql")
	err = sqlxDb.Ping()
	exception.WithError(err)

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &SqlDb{
		sql:  db,
		sqlx: sqlxDb,
	}
}

func (s *SqlDb) Query(destination interface{}, query string, params ...interface{}) {
	rows, err := s.sqlx.Queryx(query, params...)
	exception.WithError(err)
	err = rows.StructScan(destination)
	exception.WithError(err)
}

func (s *SqlDb) Exec(query string, params ...interface{}) {
	_, err := s.sqlx.Exec(query, params...)
	exception.WithError(err)
}
