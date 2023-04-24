package persistence

import (
	"fmt"

	"github.com/ruohuaii/wufumall/domain/entity"
	"github.com/ruohuaii/wufumall/domain/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/3/24
* Time: 12:08
**/

type Repositories struct {
	UserRepo          repository.UserRepository
	RequestRecordRepo repository.RequestRecordRepository
	db                *gorm.DB
	schema.NamingStrategy
}

func NewRepositories(DbUser, DbPassword, DbHost, DbPort, DbName string) (*Repositories, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUser, DbPassword, DbHost, DbPort, DbName,
	)
	config := &gorm.Config{
		NamingStrategy: &Repositories{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	}
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		UserRepo:          NewUserRepo(db),
		RequestRecordRepo: NewRequestRecordRepo(db),
		db:                db,
	}, nil
}

func (s *Repositories) AutoMigrate() error {
	return s.db.AutoMigrate(&entity.User{}, &entity.RequestRecord{})
}

func (s *Repositories) TableName(table string) string {
	return "wfm_" + s.NamingStrategy.TableName(table)
}
