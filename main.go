package main

import (
	"CleanArchitecture/config"
	"CleanArchitecture/features/users/handler"
	"CleanArchitecture/features/users/models"
	"CleanArchitecture/features/users/service"
	"CleanArchitecture/helper/jwt"
	"CleanArchitecture/routes"
	"CleanArchitecture/utils/database"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var config = config.InitConfig()

	db := database.InitDatabase(*config)
	database.Migrate(db)

	userModel := models.NewUserModels(db)
	jwtInterface := jwt.NewJWT(config.Secret, config.RefreshSecret)
	userServices := service.NewUserService(userModel, jwtInterface)

	userControll := handler.NewUserHandler(userServices)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.RouteUser(e, userControll, *config)
	e.Logger.Fatalf(e.Start(fmt.Sprintf(":%d", config.ServerPort)).Error())
}
