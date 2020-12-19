package model

// Category is artical's category
type Category struct {
	ID   int    `xorm:"pk autoincr INT(11)" json:"i_d"` // 分类标识ID
	Name string `xorm:"unique notnull" json:"name"` // 分类名
}
