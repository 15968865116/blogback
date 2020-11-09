package model

//数据库模型 用于建立数据表或者读取数据等
type User struct {
	Id   int64  `xorm:"pk autoincr INT(11)"`
	Account string `xorm:"unique notnull" json:"account"`
	Name string `xorm:"notnull" json:"name"`
	Password string `xorm:"notnull" json:"password"`
}
