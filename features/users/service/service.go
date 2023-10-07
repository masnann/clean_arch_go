package service

import (
	"CleanArchitecture/features/users"
	"CleanArchitecture/helper/jwt"
	"errors"
	"fmt"
	"strings"
)

type UserService struct {
	d users.UserModelsInterface
	j jwt.JWTInterface
}

func NewUserService(data users.UserModelsInterface, jwt jwt.JWTInterface) users.UserServiceInterface {
	return &UserService{
		d: data,
		j: jwt,
	}
}

func (us *UserService) GetAllUsers() ([]*users.User, error) {
	usersList, err := us.d.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return usersList, nil
}

func (us *UserService) Register(newData users.User) (*users.User, error) {
	result, err := us.d.Register(newData)
	if err != nil {
		return nil, errors.New("insert process failed")
	}

	return result, nil
}

func (us *UserService) Login(hp string, password string) (*users.UserCredential, error) {
	result, err := us.d.Login(hp, password)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, errors.New("data not found")
		}
		return nil, errors.New("process failed")
	}

	userIDStr := fmt.Sprint(result.ID)

	tokenData := us.j.GenerateJWT(userIDStr)

	if tokenData == nil {
		return nil, errors.New("token process failed")
	}

	response := new(users.UserCredential)
	response.Name = result.Name
	response.Access = tokenData

	return response, nil
}
