package controllers

import (
	"encoding/base64"
	"finalgo/dao"
	"finalgo/model"
	"finalgo/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
	"time"
)

type Picture struct {
	IMG string `json:"img"`
	Token string `json:"token"`
	Puberaccount string `json:"puberaccount"`
}

type Picturecontroller struct {

}

func (pc Picturecontroller)Router(engine *gin.Engine)  {
	engine.POST("/picture/blog", tool.Tokencheck, pc.Getpictureforblog)
	engine.POST("/picture/portrait",tool.Tokencheck, pc.Getpictureforuser)
}

// 文章内添加图片
func (pc Picturecontroller)Getpictureforblog(context *gin.Context)  {
	// 定义转码和路径
	var path string
	var path_forweb string
	var enc = base64.StdEncoding
	var img string

	// 接收图片对象
	var nowint = time.Now().Unix()
	var now =  strconv.FormatInt(nowint,10)
	var picture Picture
	err := context.BindJSON(&picture)
	if err!= nil {
		log.Err(err)
	}


	//为图片创造路径
	if picture.IMG[11] == 'j' {
		path = "./picturefile/blogpic/" + picture.Puberaccount+now + ".jpg"
		path_forweb =picture.Puberaccount+ now + ".jpg"
		img = picture.IMG[23:]
	} else if picture.IMG[11] == 'p' {
		path = "./picturefile/blogpic/" + picture.Puberaccount+now + ".png"
		path_forweb = picture.Puberaccount+now + ".png"
		img = picture.IMG[22:]
	} else if picture.IMG[11] == 'g' {
		path = "./picturefile/blogpic/" + picture.Puberaccount+now + ".gif"
		path_forweb = picture.Puberaccount+now + ".gif"
		img = picture.IMG[22:]
	} else {
		context.JSON(150, map[string]interface{}{
			"code":    0,
			"msg":     "失败",
			"urlpath": "",
		})
		return
	}

	// 解码图片
	data, err := enc.DecodeString(img)
	if err != nil {
		tool.LogERRAdmin("解码图片失败！")
		context.JSON(150, map[string]interface{}{
			"code":    0,
			"msg":     "失败",
			"urlpath": "",
		})
		return
	}

	// 图片写入文件
	f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		tool.LogERRAdmin("写入图片数据失败！")
		context.JSON(150, map[string]interface{}{
			"code":    0,
			"msg":     "失败",
			"urlpath": "",
		})
		return
	}
	returnpath := "http://localhost:8090/blogimg/" + path_forweb
	tool.LogINFOAdmin("写入图片数据成功！")
	context.JSON(200, map[string]interface{}{
		"code":1,
		"msg":     "成功",
		"urlpath": returnpath,
	})

}

// 更换头像
func (pc Picturecontroller)Getpictureforuser(context *gin.Context)  {
	// 定义转码和路径
	var path string
	var path_forweb string
	var enc = base64.StdEncoding
	var img string

	// 接收图片对象
	var picture Picture
	err := context.BindJSON(&picture)
	if err!= nil {
		tool.LogERRAdmin("数据绑定失败！")
		context.JSON(150, map[string]interface{}{
			"code":    0,
			"msg":     "失败",
			"urlpath": "",
		})
		return
	}

	//为图片创造路径
	//fmt.Println(picture.IMG)
	var timestring = strconv.FormatInt(time.Now().Unix(), 10)
	if picture.IMG[11] == 'j' {
		path = "./picturefile/userpic/" + picture.Puberaccount + timestring + ".jpg"
		path_forweb = picture.Puberaccount + timestring + ".jpg"
		img = picture.IMG[23:]
	} else if picture.IMG[11] == 'p' {
		path = "./picturefile/userpic/" + picture.Puberaccount + timestring + ".png"
		path_forweb = picture.Puberaccount + timestring + ".png"
		img = picture.IMG[22:]
	} else if picture.IMG[11] == 'g' {
		path = "./picturefile/userpic/" + picture.Puberaccount + timestring + ".gif"
		path_forweb = picture.Puberaccount + timestring + ".gif"
		img = picture.IMG[22:]
	} else {
		fmt.Println("buzhidhigaileix")
	}

	// 解码图片
	data, err := enc.DecodeString(img)
	if err != nil {
		tool.LogERRAdmin("解码图片失败！")
		context.JSON(150, map[string]interface{}{
			"code":    0,
			"msg":     "失败",
			"urlpath": "",
		})
	} else {
		// 图片写入文件
		f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer f.Close()
		_, err = f.Write(data)
		if err != nil {
			tool.LogERRAdmin("图片数据写入失败！")
		}
		returnpath := "http://localhost:8090/portrait/" + path_forweb
		var user = model.User{Portrait: returnpath}
		udb := dao.Userdao{tool.DBengine}
		h := udb.Updateusermessage(picture.Puberaccount, user)
		if h == 0 {
			tool.LogERRAdmin("图片更新失败！")
			context.JSON(150, map[string]interface{}{
				"code":    0,
				"msg":     "失败",
				"urlpath": "",
			})
		} else {
			tool.LogINFOAdmin("图片更新成功！")
			context.JSON(200, map[string]interface{}{
				"code":    1,
				"msg":     "成功",
				"urlpath": returnpath,
			})
		}
	}

}
