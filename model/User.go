package model

//数据库模型 用于建立数据表或者读取数据等
type User struct {
	Id        int64  `xorm:"pk autoincr INT(11)"`  // 用户ID
	Account   string `xorm:"unique notnull" json:"account"` // 账号
	Name      string `xorm:"notnull" json:"name"` // 昵称
	Password  string `xorm:"notnull" json:"password"` // 密码
	Portrait  string `xorm:"" json:"portrait"` // 头像
	Email     string `xorm:"" json:"email"` // 邮件
	Tag       string `xorm:"" json:"tag"` // 标签
	Introduce string `xorm:"" json:"introduce"` // 简介
}
