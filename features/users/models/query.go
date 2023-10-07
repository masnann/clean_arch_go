package models

import (
	"CleanArchitecture/features/users"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserModels struct {
	gorm *gorm.DB
}

func NewUserModels(g *gorm.DB) users.UserModelsInterface {
	return &UserModels{
		gorm: g,
	}
}

func (ud *UserModels) GetAllUsers() ([]*users.User, error) {
	var dbData []*User

	if err := ud.gorm.Find(&dbData).Error; err != nil {
		return nil, err
	}

	var result []*users.User
	for _, data := range dbData {
		user := &users.User{
			Name:     data.Name,
			Email:    data.Email,
			Password: data.Password,
		}
		result = append(result, user)
	}

	return result, nil
}

func (ud *UserModels) Register(newData users.User) (*users.User, error) {
	var dbData = new(User)
	dbData.Name = newData.Name
	dbData.Email = newData.Email
	dbData.Password = newData.Password

	if err := ud.gorm.Create(dbData).Error; err != nil {
		return nil, err
	}

	return &newData, nil
}

func (ud *UserModels) Login(email string, password string) (*users.User, error) {
	var dbData = new(User)
	dbData.Email = email
	dbData.Password = password

	if err := ud.gorm.Where("email = ? AND password = ?", dbData.Email, dbData.Password).First(dbData).Error; err != nil {
		logrus.Info("db error:", err.Error())
		return nil, err
	}

	var result = new(users.User)
	result.ID = dbData.ID
	result.Name = dbData.Name

	return result, nil
}
