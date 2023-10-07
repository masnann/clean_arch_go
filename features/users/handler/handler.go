package handler

import (
	"CleanArchitecture/features/users"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	s users.UserServiceInterface
}

func NewUserHandler(service users.UserServiceInterface) users.UserHandlerInterface {
	return &UserHandler{
		s: service,
	}
}
func (uh *UserHandler) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		usersList, err := uh.s.GetAllUsers()
		if err != nil {
			c.Logger().Error("handler: failed to fetch all users:", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Server Error",
			})
		}

		var response = make([]GetALLResponse, 0)
		for _, user := range usersList {
			response = append(response, GetALLResponse{
				Name:  user.Name,
				Email: user.Email,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Get All Users",
			"data":    response,
		})
	}
}

func (uh *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(CreateUserRequest)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("handler: bind input error:", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Bad Request",
			})
		}

		var serviceInput = new(users.User)
		serviceInput.Name = input.Name
		serviceInput.Email = input.Email
		serviceInput.Password = input.Password

		result, err := uh.s.Register(*serviceInput)

		if err != nil {
			c.Logger().Error("handler: input process error:", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Bad Request",
			})
		}

		var response = new(CreateResponse)
		response.Name = result.Name
		response.Email = result.Email

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Success Register",
			"data":    response,
		})
	}
}

func (uh *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginRequest)

		if err := c.Bind(input); err != nil {
			c.Logger().Error("handler: bind input error:", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Bad Request",
			})
		}

		result, err := uh.s.Login(input.Email, input.Password)

		if err != nil {
			c.Logger().Error("handler: login process error:", err.Error())
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"message": "Not Found",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Status Internal Server Error",
			})
		}

		var response = new(LoginResponse)
		response.Name = result.Name
		response.Token = result.Access

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Login",
			"data":    response,
		})
	}
}
