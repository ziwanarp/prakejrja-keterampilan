package routes

import (
	"os"
	"golang/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) *echo.Echo{
	e.Use(middleware.Logger())
	e.POST("/login", controllers.LoginController)
	
	authGroup := e.Group("")
	authGroup.Use(echojwt.JWT([]byte(os.Getenv("KEY_ENCRYPTOPN"))))
	authGroup.GET("/users", controllers.UserController)
	authGroup.GET("/user/:id", controllers.DetailUserController)
	authGroup.POST("/user", controllers.CreateUserController)

	authGroup.GET("/films", controllers.FilmController)
	authGroup.GET("/film/:id", controllers.DetailFilmController)
	authGroup.POST("/film", controllers.CreateFilmController)
	
	return e
}