package worm

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/weirwei/ikit/ilog"

	_ "github.com/mattn/go-sqlite3"
)

func TestWorm(t *testing.T) {
	orm()
}

func Source() {
	db, _ := sql.Open("sqlite3", "gee.db")
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	_, _ = db.Exec("CREATE TABLE User(Name text);")
	result, err := db.Exec("INSERT INTO User(`Name`) VALUES (?), (?)", "Tom", "Sam")
	if err == nil {
		affected, _ := result.RowsAffected()
		ilog.Info(affected)
	}
	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	var name string
	if err := row.Scan(&name); err == nil {
		ilog.Info(name)
	}
}

func orm() {
	engine, _ := NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
