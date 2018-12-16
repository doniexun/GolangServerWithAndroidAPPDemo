package models

import (
	"strings"
	"errors"
	"encoding/hex"
	"crypto/md5"
	
	"github.com/astaxie/beego/orm"
)


type User struct {
	Id       int    `orm:"column(Id)"`
	Username string `orm:"size(16);unique;column(Username)"` 
	Email    string `orm:"size(50);unique;column(Email);default();"`
	Password string `orm:"size(32);column(Password)"`
}

func NewUser() *User {
	return &User{}
}

func GetTableUser() string {
	return "user"
}

func (u *User) UserAdd(username, email, password string) (error, string, int64) {
	var (
		user User
		o    = orm.NewOrm()
	)

	l := strings.Count(username, "") - 1
	if l < 2 || l > 16 {
		return errors.New("用户名长度限制在2-16个字符"), "用户名长度限制在2-16个字符！", 0
	}
	
	if o.QueryTable(GetTableUser()).Filter("Username", username).One(&user); user.Id > 0 {
		return errors.New("用户名已被注册，请更换新的用户名"), "用户名已被注册，请更换新的用户名！", 0
	}
	
	pwd := MyMD5(password)	
	user = User{Email: email, Username: username, Password: pwd}
	uid, err := o.Insert(&user); 
	if err != nil || uid == 0 {
		return err, "用户注册失败，请更换用户名或用户邮箱后重新注册！", 0
	}
	
	return nil, "恭喜，用户注册成功！", uid
}

func MyMD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}




