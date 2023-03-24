package main

import (
	"github.com/ruohuaii/wufumall/infrastructure/persistence"
	apiv1 "github.com/ruohuaii/wufumall/interfaces/api/v1"

	"github.com/cloudwego/hertz/pkg/app/server"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/3/24
* Time: 9:44
**/

func main() {
	DbUser := "root"
	DbPassword := "dhf"
	DbHost := "127.0.0.1"
	DbPort := "3306"
	DbName := "wufumall"
	repositories, err := persistence.NewRepositories(DbUser, DbPassword, DbHost, DbPort, DbName)
	if err != nil {
		panic(err.Error())
	}

	err = repositories.AutoMigrate()
	if err != nil {
		panic(err.Error())
	}

	svr := server.New()

	v1 := svr.Group("/v1")
	{
		user := apiv1.NewUser(repositories.UserRepo)
		v1.GET("/signup", user.SignUp)
	}

	svr.Spin()
}
