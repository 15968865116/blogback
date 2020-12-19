package model

import "time"

// 博文ID、发布日期、发表用户、博文标题、博文内容、点赞数、回复数、游览量。
type Blog struct {
	ID int64 `xorm:"pk autoincr"` // 用户ID
	Pubdate time.Time `xorm:"timestamp"` // 发布日期
	Updatedate time.Time `xorm:"timestamp"` // 更新日期
	Puber string `xorm:""` // 发布者
	Puberaccount string `xorm:""` // 发布账号
	Title string `xorm:""` // blog标题
	Content string `xorm:"Text"` // blog内容
	Likecount int64  `xorm:"int default 0"` // 点赞数
	Replycount int64 `xorm:"int default 0"` // 回复数
	Scancount int64 `xorm:"int default 0"` // 浏览数
	CategoryID int `xorm:"int notnull"` // 分类ID
}

// 博文分类表
