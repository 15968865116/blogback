package controllers

import (
	"encoding/base64"
	"finalgo/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"os"
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
	engine.POST("/picture/blog",pc.Getpictureforblog)
}

// 文章内添加图片
func (pc Picturecontroller)Getpictureforblog(context *gin.Context)  {
	// 定义转码和路径
	var path string
	var path_forweb string
	var enc = base64.StdEncoding
	var img string

	// 接收图片对象
	var now = string(time.Now().Unix())
	var picture Picture
	err := context.BindJSON(&picture)
	if err!= nil {
		log.Err(err)
	}

	tokenture, err := tool.Getjwt(picture.Puberaccount,picture.Token)
	if err != nil {
		return
	}
	if tokenture != true {
		context.JSON(250,map[string]interface{}{
			"code": 0,
			"msg":"token错误或者登录过期",
		})
	} else {

		//为图片创造路径
		fmt.Println(picture.IMG)
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
			fmt.Println("buzhidhigaileix")
		}

		// 解码图片
		data, err := enc.DecodeString(img)
		if err != nil {
			log.Err(err)
		}

		// 图片写入文件
		f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer f.Close()
		_, err = f.Write(data)
		if err != nil {
			log.Err(err)
		}
		returnpath := "http://localhost:8090/blogimg/" + path_forweb
		context.JSON(200, map[string]interface{}{
			"code":1,
			"msg":     "成功",
			"urlpath": returnpath,
		})
	}
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
		log.Err(err)
	}

	tokenture, err := tool.Getjwt(picture.Puberaccount,picture.Token)
	if err != nil {
		return
	}
	if tokenture != true {
		context.JSON(250,map[string]interface{}{
			"code": 0,
			"msg":"token错误或者登录过期",
		})
	} else {

		//为图片创造路径
		fmt.Println(picture.IMG)
		if picture.IMG[11] == 'j' {
			path = "./picturefile/userpic/" + picture.Puberaccount + ".jpg"
			path_forweb =picture.Puberaccount + ".jpg"
			img = picture.IMG[23:]
		} else if picture.IMG[11] == 'p' {
			path = "./picturefile/userpic/" + picture.Puberaccount + ".png"
			path_forweb = picture.Puberaccount + ".png"
			img = picture.IMG[22:]
		} else if picture.IMG[11] == 'g' {
			path = "./picturefile/userpic/" + picture.Puberaccount + ".gif"
			path_forweb = picture.Puberaccount + ".gif"
			img = picture.IMG[22:]
		} else {
			fmt.Println("buzhidhigaileix")
		}

		// 解码图片
		data, err := enc.DecodeString(img)
		if err != nil {
			log.Err(err)
		}

		// 图片写入文件
		f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer f.Close()
		_, err = f.Write(data)
		if err != nil {
			log.Err(err)
		}
		returnpath := "http://localhost:8090/portrait/" + path_forweb
		context.JSON(200, map[string]interface{}{
			"code":1,
			"msg":     "成功",
			"urlpath": returnpath,
		})
	}
}
