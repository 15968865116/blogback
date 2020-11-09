package controllers

import (
	"finalgo/dao"
	"finalgo/model"
	"finalgo/tool"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"time"
)

// post传递过来的参数结构
type Blogpost struct {
	Puber string `json:"puber"`
	Puberaccount string `json:"puberaccount"`
	Title string `json:"title"`
	Content string `json:"content"`
	Token string `json:"token"`
}

type Blogcontroller struct {

}

func (bc Blogcontroller)Router(engine *gin.Engine)  {
	engine.POST("/blog/createblog", bc.Createnewblog)
}

// 新增 blog
func (bc Blogcontroller)Createnewblog(context *gin.Context){
	// 获得结构体
	var blogpost Blogpost
	err := context.BindJSON(&blogpost)
	if err != nil {
		log.Error().Err(err)
	}

	// 判断token情况
	tokenture, err := tool.Getjwt(blogpost.Puberaccount,blogpost.Token)
	if err != nil {
		return
	}
	if tokenture != true {
		context.JSON(250,map[string]interface{}{
			"code": 0,
			"msg":"token错误或者登录过期",
		})
	} else{
		// 执行插入数据库
		bd := dao.Blogdao{tool.DBengine}
		var blogmodel model.Blog
		blogmodel.Puberaccount = blogpost.Puberaccount
		blogmodel.Puber = blogpost.Puber
		blogmodel.Content = blogpost.Content
		blogmodel.Pubdate = time.Now()
		blogmodel.Title = blogpost.Title
		blogmodel.Updatedate = time.Now()
		result := bd.CreateBlog(blogmodel)
		if result > 0 {
			context.JSON(200,map[string]interface{}{
				"code":1,
				"msg":"添加成功",
			})
		} else {
			context.JSON(250,map[string]interface{}{
				"code":0,
				"msg":"添加失败",
			})
		}
	}
}
