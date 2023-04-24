package entity

import "time"

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/3/24
* Time: 9:13
**/

type User struct {
	ID uint64 `gorm:"primary_key;autoIncrement" json:"id"`
	//UserId 用户标记
	UserId uint64 `gorm:"not null" json:"user_id"`
	//FirstName 名
	FirstName string `gorm:"size:30;not null" json:"first_name"`
	//LastName 姓
	LastName string `gorm:"size:32;not null" json:"last_name"`
	//Phone 手机号
	Phone string `gorm:"size:32" json:"phone"`
	//Password 密码
	Password string `gorm:"size:100;not null;" json:"password"`
	//HashSalt 密码填充项
	HashSalt string `gorm:"size:32;not null" json:"hash_salt"`
	//RegistrationTime 注册时间
	RegistrationTime time.Time `gorm:"not null" json:"registration_time"`
	//Birthday 生日
	Birthday time.Time `gorm:"not null" json:"birthday"`
	//CreatedAt 数据创建时间
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	//UpdatedAt 数据最后更新时间
	UpdatedAt time.Time `gorm:"autoCreateTime" json:"updated_at"`
}
