package dao

import (
	"finalgo/model"
	"finalgo/tool"
	"github.com/rs/zerolog/log"
)

// 用于操作数据库 data access object

type Userdao struct {
	*tool.Orm
}

func (ud *Userdao)Insertuser(user model.User) int64  {
	result, err := ud.InsertOne(&user)
	if err != nil{
		log.Error().Err(err)
	}
	return result
}

// 通过用户名密码查询是否存在这个用户
func (ud *Userdao)Selectuser(account string, password string) *model.User{
	var user model.User
	b, _ := ud.Where("account = ?",account).And("password = ?",password).Get(&user)
	if b{
		return &user
	}
	return nil
}

// 通过用户名查询信息
func (ud *Userdao)SelectuserMessage(account string) *model.User{
	var user model.User
	b, _ := ud.Where("account = ?",account).Get(&user)
	if b{
		return &user
	}
	return nil
}

// 通过昵称查询信息
func (ud *Userdao)SelectuserMessageGet(name string) *model.User{
	var user model.User
	b, _ := ud.Where("name = ?",name).Get(&user)
	if b{
		return &user
	}
	return nil
}

// 更新用户信息
func (ud *Userdao)Updateusermessage(acc string,user model.User) int64{
	col,err := ud.Where("account = ?", acc).Update(&user)
	if err != nil {
		return 0
	}
	return col
}