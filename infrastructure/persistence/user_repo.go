package persistence

import (
	"github.com/ruohuaii/wufumall/domain/entity"
	"github.com/ruohuaii/wufumall/domain/repository"

	"gorm.io/gorm"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/3/24
* Time: 12:13
**/

var _ repository.UserRepository = &UserRepo{}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Save(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) GetByUserId(userid uint64) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.Where(&entity.User{UserId: userid}).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
