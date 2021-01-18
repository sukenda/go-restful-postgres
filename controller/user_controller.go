package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sukenda/go-restful-postgre/exception"
	"github.com/sukenda/go-restful-postgre/model"
	"github.com/sukenda/go-restful-postgre/service"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{
		UserService: *userService,
	}
}

func (controller *UserController) Route(app *echo.Echo) {
	app.POST("/api/signup", controller.Register)
	app.POST("/api/login", controller.Login)
}

func (controller *UserController) Register(c echo.Context) error {
	var request model.CreateUserRequest
	err := c.Bind(&request)
	exception.PanicIfNeeded(err)

	response := controller.UserService.Register(request)
	return c.JSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) Login(c echo.Context) error {
	var request model.CreateUserRequest
	err := c.Bind(&request)

	exception.PanicIfNeeded(err)

	response := controller.UserService.Login(request)
	return c.JSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
