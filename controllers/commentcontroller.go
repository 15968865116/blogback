package controllers

import (
	"finalgo/dao"
	"finalgo/model"
	"finalgo/tool"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Commentcontroller struct {

}


func (cc *Commentcontroller)Router(engine *gin.Engine)  {
	engine.POST("/comment/comment",tool.Tokencheck,cc.CreateComment)
	engine.GET("/comment/commentchecked",cc.SelectcommentChecked)
	engine.PUT("/comment/comment",tool.Tokencheck,cc.CheckComment)
	engine.DELETE("/comment/comment",tool.Tokencheck,cc.Deletecomment)
	engine.GET("/comment/commentnotchecked",tool.Tokencheck, cc.SelectcommentNotCheck)
}

// 评论
func (cc *Commentcontroller)CreateComment(context *gin.Context) {
	var comment model.Comment
	err := context.BindJSON(&comment)
	comment.IfCheck = 0
	if err != nil {
		tool.LogERRVisitor("评论数据绑定失败，" + err.Error())
	}
	comment.Commenttime = time.Now()
	cd := dao.Commentdao{tool.DBengine}
	ifsuss := cd.Insertcomment(comment)
	if ifsuss == true {
		tool.LogINFOVisitor("访客："+ comment.Commentname + "评论：" + comment.Comments + "评论成功！")
		context.JSON(200,map[string]interface{}{
			"code":1,
			"msg":"success",
			"result": "您的评论等待管理员审核。",
		})
	} else {
		tool.LogINFOVisitor("访客："+ comment.Commentname + "评论：" + comment.Comments + "评论失败！")
		context.JSON(250,map[string]interface{}{
			"code":0,
			"msg":"failed",
		})
	}
}

// 查询评论  已审核评论
func (cc *Commentcontroller)SelectcommentChecked(context *gin.Context) {
	blogid := context.Query("blogid")
	intblogid, err := strconv.Atoi(blogid)
	if err != nil {
		tool.LogERRAdmin("查询评论时，文章ID转换失败"+ err.Error())
	}
	cd := dao.Commentdao{tool.DBengine}
	comments := cd.Selectcomment(intblogid)
	context.JSON(200, map[string]interface{}{
		"code": 1,
		"msg": "ok",
		"result":comments,
	})
}

// 查询评论 未审核评论
func (cc *Commentcontroller)SelectcommentNotCheck(context *gin.Context) {
	cd := dao.Commentdao{tool.DBengine}
	comments := cd.GetAllComent()
	context.JSON(200, map[string]interface{}{
		"code": 1,
		"msg": "ok",
		"result":comments,
	})
}


// 审核评论
func (cc *Commentcontroller)CheckComment(context *gin.Context) {
	var comment model.Comment
	err := context.BindJSON(&comment)
	if err != nil {
		tool.LogERRAdmin("审核评论时，数据绑定失败"+ err.Error())
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg": "failed",
		})
		return
	}
	cd := dao.Commentdao{tool.DBengine}
	ok := cd.CheckComment(comment.Id)
	if ok {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg": "ok",
		})
	} else {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg": "failed",
		})
	}
}

// 删除评论
func (cc *Commentcontroller)Deletecomment(context *gin.Context) {
	var comment model.Comment
	err := context.BindJSON(&comment)
	if err != nil {
		tool.LogERRAdmin("审核评论时，数据绑定失败"+ err.Error())
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg": "failed",
		})
		return
	}
	cd := dao.Commentdao{tool.DBengine}
	ok := cd.Deletecomment(comment.Id)
	if ok {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg": "ok",
		})
	} else {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg": "failed",
		})
	}
}
