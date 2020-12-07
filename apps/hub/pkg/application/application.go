package application

import (
	"github.com/atastrophic/go-ms-with-eks/pkg/exception"
	"github.com/atastrophic/go-ms-with-eks/pkg/handlers"
	"github.com/atastrophic/go-ms-with-eks/pkg/repositories"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type AppDepIn interface {
	UsersHandler() *handlers.UsersHandler
	InfoHandler() *handlers.InfoHandler

	Schema() *repositories.Schema
}

type Application struct {
	builder AppDepIn
}

func NewApplication(builder AppDepIn) *Application {
	return &Application{
		builder: builder,
	}
}

func (app *Application) Start() {

	e := echo.New()
	e.Logger.SetLevel(log.OFF)

	app.setupRoutes(e)

	app.builder.Schema().Execute()

	err := e.Start(":8080")
	exception.WithError(err)
}

func (app *Application) setupRoutes(e *echo.Echo) {

	app.builder.InfoHandler().Routes(e.Group("/info"))
	app.builder.UsersHandler().Routes(e.Group("/users"))

}
