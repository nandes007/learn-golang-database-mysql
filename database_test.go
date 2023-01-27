package learn_golang_database_mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/learn_golang_database")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
