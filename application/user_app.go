package application

import (
	"github.com/ruohuaii/wufumall/domain/entity"
	"github.com/ruohuaii/wufumall/domain/repository"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/3/24
* Time: 10:15
**/

var _ UserAppInterface = &userApp{}

type userApp struct {
	repo repository.UserRepository
}

type UserAppInterface interface {
	Save(user *entity.User) error
	GetByUserName(username string) (*entity.User, error)
}

func (u *userApp) Save(user *entity.User) error {
	return u.repo.Save(user)
}

func (u *userApp) GetByUserName(username string) (*entity.User, error) {
	return u.repo.GetByUserName(username)
}
