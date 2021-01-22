package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sukenda/go-restful-postgres/config"
	"github.com/sukenda/go-restful-postgres/controller"
	"github.com/sukenda/go-restful-postgres/exception"
	"github.com/sukenda/go-restful-postgres/repository"
	"github.com/sukenda/go-restful-postgres/service"
)

func main() {

	configuration := config.New()
	database := config.NewPostgresDatabase(configuration)

	// Setup Repository
	userRepository := repository.NewUserRepository(database)
	departmentRepository := repository.NewDepartmentRepository(database)
	projectRepository := repository.NewProjectRepository(database)

	// Setup Service
	userService := service.NewUserService(&userRepository)
	departmentService := service.NewDepartmentService(&departmentRepository)
	projectService := service.NewProjectService(&projectRepository)

	// Setup Controller
	userController := controller.NewUserController(&userService)
	departmentController := controller.NewDepartmentController(&departmentService)
	projectController := controller.NewProjectController(&projectService)

	e := echo.New()

	// Error Handler
	e.HTTPErrorHandler = func(err error, context echo.Context) {
		err = exception.ErrorHandler(context, err)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	userController.Route(e)
	departmentController.Route(e)
	projectController.Route(e)

	e.Logger.Fatal(e.StartServer(config.NewEchoConfig(configuration)))
}
