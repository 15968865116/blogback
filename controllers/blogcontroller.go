package controllers

import (
	"finalgo/dao"
	"finalgo/model"
	"finalgo/tool"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// post传递过来的参数结构
type Blogpost struct {
	Puber        string `json:"puber"`
	Puberaccount string `json:"puberaccount"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Token        string `json:"token"`
	Id           int    `json:"id"`
	// CategoryID   int64    `json:"categoryid"`
	// Ifcreatenewcategory bool `json:"ifcreatenewcategory"`
	CategoryName string `json:"categoryName"`
}

// update
type Blogupdate struct {
	Puberaccount string `json:"puberaccount"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Token        string `json:"token"`
	Id           int    `json:"id"`
	CategoryName   string    `json:"categoryname"`
}

type Blogcontroller struct {
}

func (bc Blogcontroller) Router(engine *gin.Engine) {
	engine.POST("/blog/createblog", tool.Tokencheck, bc.Createnewblog)
	engine.POST("/blog/updateblog",  tool.Tokencheck, bc.Updateblog)
	engine.POST("/blog/deleteblog",  tool.Tokencheck, bc.Deleteblog)
	engine.GET("/blog/getblog", bc.Selectblog)
	engine.GET("/blog/getspecificblog", bc.Selectoneblog)
	engine.GET("/blog/blogbycategory", bc.SelectblogByCategory)
}

// 新增 blog 
func (bc Blogcontroller) Createnewblog(context *gin.Context) {
	// 获得结构体
	var blogpost Blogpost
	err := context.BindJSON(&blogpost)
	if err != nil {
		tool.LogERRAdmin("数据绑定失败" + err.Error())
	}


	// 执行插入数据库
	bd := dao.Blogdao{tool.DBengine}
	var blogmodel model.Blog
	blogmodel.Puberaccount = blogpost.Puberaccount
	blogmodel.Puber = blogpost.Puber
	blogmodel.Content = blogpost.Content
	blogmodel.Pubdate = time.Now()
	blogmodel.Title = blogpost.Title
	blogmodel.Updatedate = time.Now()

	// implement add new category
	cd := dao.Categorydao{tool.DBengine}
	ifexisted, ID := cd.SelectByName(blogpost.CategoryName)
	if ifexisted {
		blogmodel.CategoryID = ID
		result := bd.CreateBlog(blogmodel)
		if result > 0 {
			context.JSON(200, map[string]interface{}{
				"code": 1,
				"msg":  "添加成功",
			})
		} else {
			context.JSON(250, map[string]interface{}{
				"code": 0,
				"msg":  "添加失败",
			})
		}
	} else {
		var categorymodel model.Category
		categorymodel.Name = blogpost.CategoryName
		resultcate := cd.CreateCategory(categorymodel)
		if resultcate < 0 {
			tool.LogERRAdmin("新增分类失败")
			context.JSON(250, map[string]interface{}{
				"code": 0,
				"msg":  "添加失败",
			})
		} else {
			_, categoryID := cd.SelectByName(blogpost.CategoryName)
			blogmodel.CategoryID = categoryID
			result := bd.CreateBlog(blogmodel)
			if result > 0 {
				tool.LogINFOAdmin("新增文章成功！")
				context.JSON(200, map[string]interface{}{
					"code": 1,
					"msg":  "添加成功",
				})
			} else {
				tool.LogERRAdmin("新增文章失败！")
				context.JSON(250, map[string]interface{}{
					"code": 0,
					"msg":  "添加失败",
				})
			}
		}

	}
}

// update blog
func (bc Blogcontroller) Updateblog(context *gin.Context) {
	// 获得结构体
	var blogpost Blogupdate
	err := context.BindJSON(&blogpost)
	if err != nil {
		tool.LogERRAdmin("数据绑定失败:" + err.Error())
	}

	// 执行插入数据库
	bd := dao.Blogdao{tool.DBengine}
	var blogmodel model.Blog
	blogmodel.Content = blogpost.Content
	blogmodel.Title = blogpost.Title
	blogmodel.ID = int64(blogpost.Id)
	blogmodel.Updatedate = time.Now()
	cd := dao.Categorydao{tool.DBengine}
	ifexisted, ID := cd.SelectByName(blogpost.CategoryName)
	if ifexisted {
		blogmodel.CategoryID = ID
		// blogmodel.CategoryID = blogpost.CategoryID
		result := bd.UpdateBlog(blogpost.Id, blogmodel)
		if result > 0 {
			tool.LogINFOAdmin("修改成功！")
			context.JSON(200, map[string]interface{}{
				"code": 1,
				"msg":  "修改成功",
			})
		} else {
			tool.LogERRAdmin("修改失败！")
			context.JSON(250, map[string]interface{}{
				"code": 0,
				"msg":  "修改失败",
			})
		}
	} else {
		var categorymodel model.Category
		categorymodel.Name = blogpost.CategoryName
		resultcate := cd.CreateCategory(categorymodel)
		if resultcate < 0 {
			tool.LogERRAdmin("新增分类失败:")
			context.JSON(250, map[string]interface{}{
				"code": 0,
				"msg":  "添加失败",
			})
		} else {
			blogmodel.CategoryID = resultcate
			// blogmodel.CategoryID = blogpost.CategoryID
			result := bd.UpdateBlog(blogpost.Id, blogmodel)
			if result > 0 {
				tool.LogINFOAdmin("修改成功！")
				context.JSON(200, map[string]interface{}{
					"code": 1,
					"msg":  "修改成功",
				})
			} else {
				tool.LogERRAdmin("修改失败！")
				context.JSON(250, map[string]interface{}{
					"code": 0,
					"msg":  "修改失败",
				})
			}
		}
	}


}

// 得到所有的blog文章
func (bc Blogcontroller) Selectblog(context *gin.Context) {
	name := context.Query("name")
	bd := dao.Blogdao{tool.DBengine}
	result := bd.SelectBlog(name)
	tool.LogINFOAdmin("查询成功！")
	context.JSON(200, map[string]interface{}{
		"code":   1,
		"msg":    "查询成功",
		"result": result,
	})
}

// 得到某一篇文章  修改前准备
func (bc Blogcontroller) Selectoneblog(context *gin.Context) {
	id := context.Query("id")
	bd := dao.Blogdao{tool.DBengine}
	intid, err := strconv.Atoi(id)
	if err != nil {
		tool.LogERRAdmin("数据转换失败，" + err.Error())
	}
	result := bd.SelectSingleBlog(intid)
	tool.LogINFOAdmin("数据查询成功")
	context.JSON(200, map[string]interface{}{
		"code":   1,
		"msg":    "查询成功",
		"result": result,
	})
}

// 删除文章
func (bc Blogcontroller) Deleteblog(context *gin.Context) {
	// 获得结构体
	var blogpost Blogpost
	err := context.BindJSON(&blogpost)
	if err != nil {
		log.Error().Err(err)
	}

	// 判断token情况

	// 执行插入数据库
	bd := dao.Blogdao{tool.DBengine}
	result := bd.DeleteBlog(blogpost.Id)
	if result == true {
		tool.LogINFOAdmin("删除文章成功！")
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  "删除成功",
		})
	} else {
		tool.LogERRAdmin("删除文章失败！")
		context.JSON(250, map[string]interface{}{
			"code": 0,
			"msg":  "删除失败",
		})
	}

}

// 通过文章分类查询文章
func (bc Blogcontroller) SelectblogByCategory(context *gin.Context) {
	var blogs []model.Blog
	categoryid := context.Query("categoryid")
	categoryidint, err := strconv.Atoi(categoryid)
	if err != nil {
		tool.LogERRVisitor("转换失败，"+ err.Error())
		context.JSON(404, map[string]interface{}{
			"code":0,
			"msg":"获取失败",
		})
		return
	}

	bd := dao.Blogdao{tool.DBengine}
	blogs = bd.SelectBlogByCategory(categoryidint)
	context.JSON(200, map[string]interface{}{
		"code":1,
		"msg":"获取成功",
		"result": blogs,
	})
}
