package users

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

type UserCredential struct {
	Name   string
	Access map[string]any
}

type UserHandlerInterface interface {
	GetAllUsers() echo.HandlerFunc
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}
type UserServiceInterface interface {
	GetAllUsers() ([]*User, error)
	Register(newData User) (*User, error)
	Login(hp string, password string) (*UserCredential, error)
}

type UserModelsInterface interface {
	GetAllUsers() ([]*User, error)
	Register(newData User) (*User, error)
	Login(hp string, password string) (*User, error)
}
