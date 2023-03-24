package v1

import (
	"context"
	"github.com/ruohuaii/wufumall/interfaces/api"

	"github.com/ruohuaii/wufumall/application/service"
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
	userApp service.UserAppInterface
	api.Wrapper
}

func NewUser(userApp service.UserAppInterface) *User {
	return &User{userApp: userApp}
}

func (u *User) SignUp(c context.Context, ctx *app.RequestContext) (any, error) {
	request := model.SignUpRequest{}
	err := ctx.BindAndValidate(&request)
	if err != nil {
		return nil, err
	}

	type Output struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	return &Output{Name: "hhs", Age: 18}, nil
}

func (u *User) SignIn() {

}
