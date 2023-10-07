package routes

import (
	"CleanArchitecture/config"
	"CleanArchitecture/features/users"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo, uc users.UserHandlerInterface, cfg config.Config) {
	e.POST("/users", uc.Register())
	e.POST("/login", uc.Login())
	e.GET("/users", uc.GetAllUsers())
}
