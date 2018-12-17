package controllers

import (
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
	
	err, msg, uid := models.NewUser().UserAdd(username, email, password)
	if err == nil {
		// 向客户端发送 json 数据
		info := "{\"status\":200,\"msg\":" + msg + ",\"extra\":" + strconv.FormatInt(uid, 10) + "}"
		c.Ctx.Output.Body([]byte(info))
	} else {
		info := "{\"status\":500,\"msg\":" + msg + ",\"extra\":" + strconv.FormatInt(uid, 10) + "}"
		c.Ctx.Output.Body([]byte(info))
	}
}

