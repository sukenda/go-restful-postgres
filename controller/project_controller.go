package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sukenda/go-restful-postgres/exception"
	"github.com/sukenda/go-restful-postgres/model"
	"github.com/sukenda/go-restful-postgres/service"
	"net/http"
)

type ProjectController struct {
	ProjectService service.ProjectService
}

func NewProjectController(departmentService *service.ProjectService) ProjectController {
	return ProjectController{
		ProjectService: *departmentService,
	}
}

func (controller *ProjectController) Route(app *echo.Echo) {
	app.POST("/api/projects", controller.Insert)
	app.GET("/api/projects", controller.Find)
}

func (controller *ProjectController) Insert(c echo.Context) error {
	var request model.CreateProjectRequest
	err := c.Bind(&request)
	exception.PanicIfNeeded(err)

	response := controller.ProjectService.Insert(request)
	return c.JSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProjectController) Find(c echo.Context) error {
	response := controller.ProjectService.Find()
	return c.JSON(http.StatusOK, model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
