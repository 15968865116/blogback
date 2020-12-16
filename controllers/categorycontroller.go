package controllers

import (
	"finalgo/dao"
	"finalgo/model"
	"finalgo/tool"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Categorycontroller is the controller of category
type Categorycontroller struct {
}

// Router is the router of the controller
func (ca *Categorycontroller) Router(engine *gin.Engine) {
	engine.GET("category/getallcategory", ca.Getallcategory)
	engine.DELETE("category/deletecategory", ca.Deletecategory)
	engine.PUT("category/updatecategory", ca.Updatecategory)
}

// Getallcategory is Get all category
func (ca *Categorycontroller) Getallcategory(context *gin.Context) {
	cd := dao.Categorydao{tool.DBengine}
	CategoryList := cd.Selectallcategory()
	context.JSON(200, map[interface{}]interface{}{
		"code":     1,
		"status":   "success",
		"category": CategoryList,
	})
}

// Deletecategory is delete category
func (ca *Categorycontroller) Deletecategory(context *gin.Context) {
	cd := dao.Categorydao{tool.DBengine}
	var modelcategory model.Category
	err := context.BindJSON(&modelcategory)
	if err != nil {
		log.Error().Err(err)
	}
	result := cd.Deletecategorybycategory(modelcategory)

	if result {
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
	if err != nil {
		log.Error().Err(err)
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
