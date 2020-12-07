package source

import (
	"github.com/atastrophic/go-ms-with-eks/pkg/exception"
	"github.com/atastrophic/go-ms-with-eks/pkg/sdb"
	"github.com/gchaincl/dotsql"
)

type Source interface {
	Query(object interface{}, squery string, params ...interface{})
	Exec(query string, params ...interface{})
}

type _SqlSource struct {
	sqldot *dotsql.DotSql
	sqldb  *sdb.SqlDb
}

func NewSqlSource(sqldot *dotsql.DotSql, sql *sdb.SqlDb) Source {
	return &_SqlSource{
		sqldot: sqldot,
		sqldb:  sql,
	}
}

func (s *_SqlSource) Query(object interface{}, query string, params ...interface{}) {
	q, err := s.sqldot.Raw(query)
	exception.WithError(err)
	s.sqldb.Query(object, q, params...)
}

func (s *_SqlSource) Exec(query string, params ...interface{}) {
	q, err := s.sqldot.Raw(query)
	exception.WithError(err)
	s.sqldb.Exec(q, params...)
}
