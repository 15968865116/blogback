package controllers

import (
	"finalgo/dao"
	"finalgo/model"
	"finalgo/tool"
	"github.com/gin-gonic/gin"
)

// Categorycontroller is the controller of category
type Categorycontroller struct {
}

type CategoryDelete struct {
	ID   int    `json:"i_d"` // 分类标识ID
	Name string `json:"name"` // 分类名
	IfdeleteArticle bool `json:"ifdelete_article"`
}

// Router is the router of the controller
func (ca *Categorycontroller) Router(engine *gin.Engine) {
	engine.GET("category/allcategoryAdimin", ca.GetallcategoryAdmin)
	engine.GET("category/allcategoryVisitor", ca.GetallcategoryVisitor)
	engine.DELETE("category/category", tool.Tokencheck, ca.Deletecategory)
	engine.PUT("category/category",tool.Tokencheck, ca.Updatecategory)
}

// Getallcategory is Get all category
func (ca *Categorycontroller) GetallcategoryAdmin(context *gin.Context) {
	cd := dao.Categorydao{tool.DBengine}
	CategoryList := cd.Selectallcategory()
	tool.LogINFOAdmin("获取所有分类信息成功！")
	context.JSON(200, map[string]interface{}{
		"code":     1,
		"status":   "success",
		"category": CategoryList,
	})
}


// Getallcategory is Get all category
func (ca *Categorycontroller) GetallcategoryVisitor(context *gin.Context) {
	cd := dao.Categorydao{tool.DBengine}
	CategoryList := cd.Selectallcategory()
	tool.LogINFOVisitor("获取所有分类信息成功！")
	context.JSON(200, map[string]interface{}{
		"code":     1,
		"status":   "success",
		"category": CategoryList,
	})
}

// Deletecategory is delete category
func (ca *Categorycontroller) Deletecategory(context *gin.Context) {
	cd := dao.Categorydao{tool.DBengine}
	var modelcategory model.Category
	var modelcategorydelete CategoryDelete
	err := context.BindJSON(&modelcategorydelete)
	if err != nil {
		tool.LogERRAdmin("绑定分类删除请求数据失败，" + err.Error())
		context.JSON(300, map[string]interface{}{
			"code":   0,
			"delete": "failed",
			"msg":"解析参数失败",
		})
		return
	}
	modelcategory.ID = modelcategorydelete.ID
	modelcategory.Name = modelcategorydelete.Name

	// fmt.Println(modelcategorydelete.IfdeleteArticle)
	result := cd.Deletecategorybycategory(modelcategory)

	if result {
		if modelcategorydelete.IfdeleteArticle {
			bd := dao.Blogdao{tool.DBengine}
			true := bd.DeleteByCategory(modelcategory.ID)
			if true {
				tool.LogINFOAdmin("删除相应分类文章数据成功！")
			} else {
				tool.LogINFOAdmin("删除相应分类文章数据失败！")
			}

		}
		tool.LogINFOAdmin("删除文章分类数据成功")
		context.JSON(200, map[string]interface{}{
			"code":   1,
			"delete": "success",
		})
	} else {
		tool.LogERRAdmin("删除文章分类数据失败！")
		context.JSON(300, map[string]interface{}{
			"code":   0,
			"delete": "failed",
		})
	}
}

// Updatecategory is update category
func (ca *Categorycontroller) Updatecategory(context *gin.Context) {
	cd := dao.Categorydao{tool.DBengine}
	var modelcategory model.Category
	err := context.BindJSON(&modelcategory)
	if err != nil {
		tool.LogERRAdmin("绑定文章分类数据失败！")
		context.JSON(300, map[string]interface{}{
			"code":   0,
			"update": "failed",
		})
		return
	}
	result := cd.Updatecategory(modelcategory)
	if result {
		tool.LogINFOAdmin("更新文章分类成功！")
		context.JSON(200, map[string]interface{}{
			"code":   1,
			"update": "success",
		})
	} else {
		tool.LogERRAdmin("更新文章分类失败！")
		context.JSON(300, map[string]interface{}{
			"code":   0,
			"update": "failed",
		})
	}
}

// create category put in create blog ?
