package controllers

import (
	"finalgo/dao"
	"finalgo/model"
	"finalgo/tool"
	"fmt"
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
	engine.GET("category/getallcategory", ca.Getallcategory)
	engine.DELETE("category/deletecategory", tool.Tokencheck, ca.Deletecategory)
	engine.PUT("category/updatecategory",tool.Tokencheck, ca.Updatecategory)
}

// Getallcategory is Get all category
func (ca *Categorycontroller) Getallcategory(context *gin.Context) {
	cd := dao.Categorydao{tool.DBengine}
	CategoryList := cd.Selectallcategory()
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
			bd.DeleteByCategory(modelcategory.ID)
		}
		context.JSON(200, map[string]interface{}{
			"code":   1,
			"delete": "success",
		})
	} else {
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
	fmt.Println(modelcategory)
	if err != nil {
		context.JSON(300, map[string]interface{}{
			"code":   0,
			"update": "failed",
		})
		return
	}
	result := cd.Updatecategory(modelcategory)
	if result {
		context.JSON(200, map[string]interface{}{
			"code":   1,
			"update": "success",
		})
	} else {
		context.JSON(300, map[string]interface{}{
			"code":   0,
			"update": "failed",
		})
	}
}

// create category put in create blog ?
