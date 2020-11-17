package dao

import (
	"finalgo/model"
	"finalgo/tool"
	"github.com/rs/zerolog/log"
)

type Commentdao struct {
	*tool.Orm
}

// 评论
func (cd *Commentdao)Insertcomment(comment model.Comment) bool {
	result, err := cd.InsertOne(comment)
	if err != nil {
		return false
	}
	if result > 0 {
		return true
	} else {
		return false
	}
}

// 查询评论
func (cd *Commentdao)Selectcomment(blogid int) []model.Comment{
	var comments = make([]model.Comment,0)
	err := cd.Where("blogid = ?", blogid).Find(&comments)
	if err != nil {
		log.Err(err)
	}
	return comments
}