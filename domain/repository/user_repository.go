package repository

import "github.com/ruohuaii/wufumall/domain/entity"

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/3/24
* Time: 9:14
**/

type UserRepository interface {
	Save(user *entity.User) error
	GetByUserName(username string) (*entity.User, error)
}
