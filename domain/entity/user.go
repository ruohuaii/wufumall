package entity

import "time"

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/3/24
* Time: 9:13
**/

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserName  string    `gorm:"size:32;not null" json:"username"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoCreateTime" json:"updated_at"`
}
