package dao

import (
	"finalgo/model"
	"finalgo/tool"

	"github.com/rs/zerolog/log"
)

// blog 的一些数据库操作

type Blogdao struct {
	*tool.Orm
}

//插入博文
func (bd Blogdao) CreateBlog(blog model.Blog) int64 {
	result, err := bd.InsertOne(&blog)
	if err != nil {
		log.Error().Err(err)
	}
	return result
}

//更新博文
func (bd Blogdao) UpdateBlog(id int, blog model.Blog) int64 {
	affected, err := bd.Id(id).Cols("content", "updatedate", "title", "categoryid").Update(&blog)
	if err != nil {
		log.Error().Err(err)
	}
	return affected

}

//查询所有的博文
func (bd Blogdao) SelectBlog(name string) []model.Blog {
	var blogs = make([]model.Blog, 0)
	err := bd.SQL("select * from blog where puber = ?", name).Find(&blogs)
	if err != nil {
		log.Error().Err(err)
	}
	return blogs
}

// 删除相应种类的所有文章
func (bd Blogdao) DeleteByCategory(categoryID int)(bool) {
	var blog model.Blog
	blog.CategoryID = categoryID
	_, err := bd.Exec("delete from blog where category_i_d = ?", categoryID)

	if err != nil {
		return false
	}

	return true
}

// 查询单条博文
func (bd Blogdao) SelectSingleBlog(id int) model.Blog {
	var blog = model.Blog{}
	_, err := bd.SQL("select * from blog where i_d = ?", id).Get(&blog)
	if err != nil {
		log.Error().Err(err)
	}
	return blog
}

// 删除博文
func (bd Blogdao) DeleteBlog(id int) bool {
	var blog = new(model.Blog)
	b, err := bd.Id(id).Delete(blog)
	if err != nil {
		log.Error().Err(err)
		return false
	}
	if b > 0 {
		return true
	} else {
		return false
	}

}
