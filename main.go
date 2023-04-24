package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/requestid"
	"github.com/ruohuaii/wufumall/infrastructure/persistence"
	"github.com/ruohuaii/wufumall/interfaces/api/v1"
	"github.com/ruohuaii/wufumall/interfaces/middleware"
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
	repos, err := persistence.NewRepositories(DbUser, DbPassword, DbHost, DbPort, DbName)
	if err != nil {
		panic(err.Error())
	}

	err = repos.AutoMigrate()
	if err != nil {
		panic(err.Error())
	}

	svr := server.New()
	svr.Use(requestid.New())
	svr.Use(middleware.RequestTrackingMiddleware(repos.RequestRecordRepo))
	apiV1 := svr.Group("/api/v1")
	{
		user := v1.NewUser(repos.UserRepo)
		apiV1.GET("/signup", middleware.Wrap(user.SignUp))
	}

	svr.Spin()
}
