package application

import (
	"github.com/ruohuaii/wufumall/domain/entity"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/3/24
* Time: 10:15
**/

type UserAppInterface interface {
	Save(user *entity.User) error
	GetByUserId(userid uint64) (*entity.User, error)
}
