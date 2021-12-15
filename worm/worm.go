package worm

import (
	"database/sql"

	"github.com/weirwei/ikit/ilog"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		ilog.Error(err)
		return
	}
	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		ilog.Error(err)
		return
	}
	e = &Engine{db: db}
	ilog.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		ilog.Error("Failed to close database")
	}
	ilog.Info("Close database success")
}

func (engine *Engine) NewSession() *Session {
	return New(engine.db)
}
