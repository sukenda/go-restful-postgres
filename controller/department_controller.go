package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sukenda/go-restful-postgres/exception"
	"github.com/sukenda/go-restful-postgres/model"
	"github.com/sukenda/go-restful-postgres/service"
	"net/http"
)

type DepartmentController struct {
	DepartmentService service.DepartmentService
}

func NewDepartmentController(departmentService *service.DepartmentService) DepartmentController {
	return DepartmentController{
		DepartmentService: *departmentService,
	}
}

func (controller *DepartmentController) Route(app *echo.Echo) {
	app.POST("/api/departments", controller.Insert)
	app.GET("/api/departments", controller.Find)
}

func (controller *DepartmentController) Insert(c echo.Context) error {
	var request model.CreateDepartmentRequest
	err := c.Bind(&request)
	exception.PanicIfNeeded(err)

	response := controller.DepartmentService.Insert(request)
	return c.JSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *DepartmentController) Find(c echo.Context) error {
	response := controller.DepartmentService.Find()
	return c.JSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
