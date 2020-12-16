package model

import "time"

// 博文ID、发布日期、发表用户、博文标题、博文内容、点赞数、回复数、游览量。
type Blog struct {
	ID int64 `xorm:"pk autoincr"`
	Pubdate time.Time `xorm:"timestamp"`
	Updatedate time.Time `xorm:"timestamp"`
	Puber string `xorm:""`
	Puberaccount string `xorm:""`
	Title string `xorm:""`
	Content string `xorm:"Text"`
	Likecount int64  `xorm:"int default 0"`
	Replycount int64 `xorm:"int default 0"`
	Scancount int64 `xorm:"int default 0"`
	CategoryID int64 `xorm:"int notnull"`
}

// 博文分类表
