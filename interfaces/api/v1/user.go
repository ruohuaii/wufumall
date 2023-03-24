package v1

import (
	"context"

	"github.com/ruohuaii/wufumall/application"
	model "github.com/ruohuaii/wufumall/interfaces/model/user"

	"github.com/cloudwego/hertz/pkg/app"
)

/**
* Created by : GoLand
* User: self-denial
* Date: 2023/3/24
* Time: 9:20
**/

type User struct {
	userApp application.UserAppInterface
}

func NewUser(userApp application.UserAppInterface) *User {
	return &User{userApp: userApp}
}

func (u *User) SignUp(c context.Context, ctx *app.RequestContext) {
	request := model.SignUpRequest{}
	err := ctx.BindAndValidate(&request)
	if err != nil {
		ctx.JSON(200, err.Error())
		return
	}

	ctx.JSON(200, map[string]string{
		"name": "hhs",
		"msg":  "注册成功",
	})
}

func (u *User) SignIn() {

}
