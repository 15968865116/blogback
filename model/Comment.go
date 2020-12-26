package model

import "time"

type Comment struct {
	Id int `xorm:"pk autoincr INT(11)" json:"id"` // 评论ID
	Commentname string `xorm:"notnull" json:"commentname"` // 评论者姓名
	Comments string `xorm:"notnull" json:"comments"` // 评论内容
	Blogid int `xorm:"notnull" json:"blogid"` // 评论的博文ID
	Commenttime time.Time `xorm:"timestamp notnull" json:"commenttime"` // 评论的时间
	IfCheck int `xorm:"int" json:"if_check"` // 是否经过管理员审核
}