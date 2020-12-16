package model

// Category is artical's category
type Category struct {
	ID   int    `xorm:"pk autoincr INT(11)"`
	Name string `xorm:"unique notnull" json:"name"`
}
