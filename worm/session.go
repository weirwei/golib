package worm

import (
	"database/sql"
	"strings"

	"github.com/weirwei/ikit/ilog"
)

type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVars []interface{}
}

func New(db *sql.DB) *Session {
	return &Session{db: db}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

// DB 获取db
func (s *Session) DB() *sql.DB {
	return s.db
}

// Raw SQL拼接
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// Exec 执行
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	ilog.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		ilog.Error(err)
	}
	return
}

// QueryRow 单行查询
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	ilog.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

// QueryRows 多行查询
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	ilog.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		ilog.Error(err)
	}
	return
}
