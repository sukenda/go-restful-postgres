package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sukenda/go-restful-postgre/config"
	"github.com/sukenda/go-restful-postgre/controller"
	"github.com/sukenda/go-restful-postgre/repository"
	"github.com/sukenda/go-restful-postgre/service"
)

func main() {

	configuration := config.New()
	database := config.NewPostgresDatabase(configuration)

	// Setup Repository
	userRepository := repository.NewUserRepository(database)

	// Setup Service
	userService := service.NewUserService(&userRepository)

	// Setup Controller
	userController := controller.NewUserController(&userService)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	userController.Route(e)

	e.Logger.Fatal(e.StartServer(config.NewEchoConfig(configuration)))
}
