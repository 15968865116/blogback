package controllers

import (
	"finalgo/dao"
	"finalgo/model"
	"finalgo/tool"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

type Commentcontroller struct {

}

func (cc *Commentcontroller)Router(engine *gin.Engine)  {
	engine.POST("/comment/createcomment",cc.CreateComment)
	engine.GET("/comment/selectcomment",cc.Selectcomment)
}

// 评论
func (cc *Commentcontroller)CreateComment(context *gin.Context) {
	var comment model.Comment
	err := context.BindJSON(&comment)
	if err != nil {
		log.Err(err)
	}
	comment.Commenttime = time.Now()
	cd := dao.Commentdao{tool.DBengine}
	ifsuss := cd.Insertcomment(comment)
	if ifsuss == true {
		context.JSON(200,map[string]interface{}{
			"code":1,
			"msg":"success",
			"result": comment,
		})
	} else {
		context.JSON(250,map[string]interface{}{
			"code":0,
			"msg":"failed",
		})
	}
}

// 查询评论
func (cc *Commentcontroller)Selectcomment(context *gin.Context) {
	blogid := context.Query("blogid")
	intblogid, err := strconv.Atoi(blogid)
	if err != nil {
		log.Err(err)
	}
	cd := dao.Commentdao{tool.DBengine}
	comments := cd.Selectcomment(intblogid)
	context.JSON(200, map[string]interface{}{
		"code": 1,
		"msg": "ok",
		"result":comments,
	})
}
