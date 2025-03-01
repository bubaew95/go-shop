package repository

import (
	"database/sql"
	"github.com/bubaew95/go_shop/conf"
	"time"
)

type DataBase struct {
	*sql.DB
}

func NewDB(c *conf.DatabaseConfig) (*DataBase, error) {
	db, err := newDatabase(c)
	if err != nil {
		return nil, err
	}

	return &DataBase{
		db,
	}, nil
}

func newDatabase(c *conf.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open(c.Driver, c.Dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * time.Duration(c.ConnMaxLifetimeInMinute))
	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (d DataBase) Close() error {
	return d.DB.Close()
}

func (d DataBase) GetDB() *sql.DB {
	return d.DB
}
