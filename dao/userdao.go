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