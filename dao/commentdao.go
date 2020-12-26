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
	ifcheck := 0
	err := cd.Where("blogid = ?", blogid).And("if_check = ?",ifcheck).Find(&comments)
	if err != nil {
		log.Err(err)
	}
	return comments
}

// 获取所有未审查评论
func (cd *Commentdao) GetAllComent() []model.Comment {
	var comments = make([]model.Comment, 0)
	ifcheck := 0
	err := cd.Where("if_check = ?",ifcheck).Find(&comments)
	if err != nil {
		tool.LogINFOAdmin("查询评论失败")
	}
	return comments
}

// 评论审查更新
func (cd *Commentdao) CheckComment(commentid int) bool {
	_, err := cd.Exec("update comment set if_check = 1 where id = ?", commentid)
	if err != nil {
		tool.LogERRAdmin("评论数据审查更新失败。")
		return false
	}
	return true
}

// 删除违规评论
func (cd *Commentdao) Deletecomment(commentid int) bool {
	_, err := cd.Exec("delete from comment where id = ?", commentid)
	if err != nil {
		tool.LogERRAdmin("评论数据审查更新失败。")
		return false
	}
	return true
}