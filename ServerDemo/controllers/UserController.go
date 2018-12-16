package controllers

import (
	"fmt"
	"strconv"
	
	"github.com/doniexun/GolangServerWithAndroidAPPDemo/ServerDemo/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Register() {
	username := c.GetString("username")
	email := c.GetString("email")
	password := c.GetString("password")
	
	fmt.Println(username + "\n" + email + "\n" + password)
	
	_, msg, uid := models.NewUser().UserAdd(username, email, password)
	// 由于无论正确与否，注册时的提示信息都会传入 msg 变量中，故无需处理 err
	c.Ctx.Output.Body([]byte(msg + "uid is " + strconv.FormatInt(uid, 10)))
}
