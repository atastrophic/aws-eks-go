package application

import (
	"github.com/atastrophic/go-ms-with-eks/pkg/appconfig"
	"github.com/atastrophic/go-ms-with-eks/pkg/exception"
	"github.com/atastrophic/go-ms-with-eks/pkg/handlers"
	"github.com/atastrophic/go-ms-with-eks/pkg/password"
	"github.com/atastrophic/go-ms-with-eks/pkg/repositories"
	"github.com/atastrophic/go-ms-with-eks/pkg/sdb"
	"github.com/atastrophic/go-ms-with-eks/pkg/services"
	"github.com/atastrophic/go-ms-with-eks/pkg/source"
	"github.com/gchaincl/dotsql"
)

type AppDep struct {
	AppDepIn
	userHandler *handlers.UsersHandler
	infoHandler *handlers.InfoHandler
	schema      *repositories.Schema
	sdb         *sdb.SqlDb
}

func NewAppDep() *AppDep {
	return &AppDep{}
}

func (b *AppDep) UsersHandler() *handlers.UsersHandler {
	if b.userHandler == nil {
		b.userHandler = handlers.NewUsersHandler(
			services.NewUserService(
				password.NewPasswordGenerator(),
			),
		)
	}
	return b.userHandler
}

func (b *AppDep) InfoHandler() *handlers.InfoHandler {
	if b.infoHandler == nil {
		b.infoHandler = &handlers.InfoHandler{}
	}
	return b.infoHandler
}

func (b *AppDep) Schema() *repositories.Schema {
	if b.schema == nil {
		sqldot, err := dotsql.LoadFromFile("./sql/schema.sql")
		exception.WithError(err)
		b.schema = repositories.NewSchema(
			source.NewSqlSource(sqldot, b.sqldb()),
		)
	}
	return b.schema
}

func (b *AppDep) sqldb() *sdb.SqlDb {
	if b.sdb == nil {
		config := appconfig.NewConfig("default", "local")
		b.sdb = sdb.NewSDB(config.SQL)
	}

	return b.sdb
}
