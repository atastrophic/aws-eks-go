package handlers

import (
	"net/http"

	"github.com/atastrophic/go-ms-with-eks/pkg/models"
	"github.com/atastrophic/go-ms-with-eks/pkg/services"
	"github.com/labstack/echo"
)

type UsersHandler struct {
	service *services.UserService
}

func NewUsersHandler(service *services.UserService) *UsersHandler {
	return &UsersHandler{
		service: service,
	}
}

func (h *UsersHandler) Routes(router *echo.Group) {

	router.POST("/signup", h.Signup)
	router.POST("/login", h.Login)

}

// Signup Handle user logins
func (h *UsersHandler) Signup(context echo.Context) error {

	var user models.User

	err := func() error {
		err := context.Bind(&user)
		if err != nil {
			return err
		}

		err = h.service.Signup(user)
		if err != nil {
			return err
		}
		return nil
	}()

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}

	return context.JSON(http.StatusOK, user)
}

// Login handles user logins
func (h *UsersHandler) Login(context echo.Context) error {
	var user models.User
	var session *models.Session
	err := func() error {
		err := context.Bind(&user)
		if err != nil {
			return err
		}

		session, err = h.service.Login(user)
		if err != nil {
			return err
		}
		return nil
	}()

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	}

	return context.JSON(http.StatusOK, *session)
}
