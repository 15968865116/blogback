package model

import "time"

type Comment struct {
	Id int `xorm:"pk autoincr INT(11)" json:"id"`
	Commentname string `xorm:"notnull" json:"commentname"`
	Comments string `xorm:"notnull" json:"comments"`
	Blogid int `xorm:"notnull" json:"blogid"`
	Commenttime time.Time `xorm:"timestamp notnull" json:"commenttime"`
}