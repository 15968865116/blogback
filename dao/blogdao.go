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
func (bd Blogdao)CreateBlog(blog model.Blog)  int64{
	result, err := bd.InsertOne(&blog)
	if err != nil{
		log.Error().Err(err)
	}
	return result
}
