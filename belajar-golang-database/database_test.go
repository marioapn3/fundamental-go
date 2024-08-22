package belajargolangdatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/belajar_golang_database")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
