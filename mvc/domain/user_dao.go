package domain

import (
	"fmt"
	"github.com/enginbulut/golang-microservices/mvc/utils"
	"net/http"
)

var(
	users = map[int64]*User {
		123: {Id: 123, FirstName: "Engin", LastName: "Bulut", Email: "test@email.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v was not found\n", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}

	return user, nil
}
