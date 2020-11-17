package controllers

import (
	"finalgo/dao"
	"finalgo/model"
	"finalgo/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

// post传递过来的参数结构
type Blogpost struct {
	Puber string `json:"puber"`
	Puberaccount string `json:"puberaccount"`
	Title string `json:"title"`
	Content string `json:"content"`
	Token string `json:"token"`
	Id int `json:"id"`
}

// update
type Blogupdate struct {
	Puberaccount string `json:"puberaccount"`
	Title string `json:"title"`
	Content string `json:"content"`
	Token string `json:"token"`
	Id int `json:"id"`
}

type Blogcontroller struct {

}

func (bc Blogcontroller)Router(engine *gin.Engine)  {
	engine.POST("/blog/createblog", bc.Createnewblog)
	engine.POST("/blog/updateblog", bc.Updateblog)
	engine.POST("/blog/deleteblog",bc.Deleteblog)
	engine.GET("/blog/getblog",bc.Selectblog)
	engine.GET("/blog/getspecificblog",bc.Selectoneblog)
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

// update blog
func (bc Blogcontroller)Updateblog(context *gin.Context)  {
	// 获得结构体
	var blogpost Blogupdate
	err := context.BindJSON(&blogpost)
	if err != nil {
		log.Error().Err(err)
	}
	fmt.Println(blogpost)
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
		blogmodel.Content = blogpost.Content
		blogmodel.Title = blogpost.Title
		blogmodel.ID = int64(blogpost.Id)
		blogmodel.Updatedate = time.Now()
		result := bd.UpdateBlog(blogpost.Id,blogmodel)
		if result > 0 {
			context.JSON(200,map[string]interface{}{
				"code":1,
				"msg":"修改成功",
			})
		} else {
			context.JSON(250,map[string]interface{}{
				"code":0,
				"msg":"修改失败",
			})
		}
	}
}

// 得到所有的blog文章
func (bc Blogcontroller)Selectblog(context *gin.Context) {
	name := context.Query("name")
	bd := dao.Blogdao{tool.DBengine}
	result := bd.SelectBlog(name)
	context.JSON(200,map[string]interface{}{
		"code":1,
		"msg":"查询成功",
		"result":result,
	})
}
// 得到某一篇文章  修改前准备
func (bc Blogcontroller)Selectoneblog(context *gin.Context) {
	id := context.Query("id")
	bd := dao.Blogdao{tool.DBengine}
	intid, err := strconv.Atoi(id)
	if err != nil {
		log.Err(err)
	}
	result := bd.SelectSingleBlog(intid)
	context.JSON(200,map[string]interface{}{
		"code":1,
		"msg":"查询成功",
		"result":result,
	})
}

// 删除文章
func (bc Blogcontroller)Deleteblog(context *gin.Context)  {
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
		result := bd.DeleteBlog(blogpost.Id)
		if result == true {
			context.JSON(200,map[string]interface{}{
				"code":1,
				"msg":"删除成功",
			})
		} else {
			context.JSON(250,map[string]interface{}{
				"code":0,
				"msg":"删除失败",
			})
		}
	}
}
