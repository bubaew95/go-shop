package ports

import (
	"database/sql"
	"io"
)

type DataBase interface {
	io.Closer
	GetDB() *sql.DB
}
