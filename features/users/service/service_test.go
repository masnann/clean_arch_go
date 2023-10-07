package service

import (
	"CleanArchitecture/features/users"
	"CleanArchitecture/features/users/mocks"
	helper "CleanArchitecture/helper/mocks"
	"github.com/stretchr/testify/mock"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	jwt := helper.NewJWTInterface(t)
	data := mocks.NewUserModelInterface(t)
	service := NewUserService(data, jwt)
	newUser := users.User{
		Name:     "supri",
		Email:    "supri@mail.com",
		Password: "admin123",
	}
	t.Run("Success insert", func(t *testing.T) {
		data.On("Register", newUser).Return(&newUser, nil).Once()
		result, err := service.Register(newUser)
		assert.Nil(t, err)
		assert.Equal(t, newUser.ID, result.ID)
		assert.Equal(t, newUser.Name, result.Name)
		data.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	j := helper.NewJWTInterface(t)
	data := mocks.NewUserModelInterface(t)
	service := NewUserService(data, j)
	userData := users.User{
		Name:     "supri",
		Email:    "supri@mail.com",
		Password: "admin123",
	}

	t.Run("success login", func(t *testing.T) {
		jwtResult := map[string]any{"access_token": "randomAccessToken"}
		data.On("Login", userData.Email, userData.Password).Return(&userData, nil)
		j.On("GenerateJWT", mock.Anything).Return(jwtResult)
		result, err := service.Login(userData.Email, userData.Password)

		data.AssertExpectations(t)
		j.AssertExpectations(t)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "supri", result.Name)
		assert.Equal(t, jwtResult, result.Access)
	})

}

func TestGetAllUsers(t *testing.T) {
	j := helper.NewJWTInterface(t)
	data := mocks.NewUserModelInterface(t)
	service := NewUserService(data, j)
	fakeUsers := []*users.User{
		&users.User{Name: "User 1", Email: "user1@example.com"},
		&users.User{Name: "User 2", Email: "user2@example.com"},
	}

	data.On("GetAllUsers").Return(fakeUsers, nil).Once()
	result, err := service.GetAllUsers()

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, fakeUsers, result)
	data.AssertExpectations(t)
}
