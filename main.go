package main

import (
	"finalgo/controllers"
	"finalgo/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main()  {
	
	config, err := tool.ParseConfig("./config/config.json")
	if err != nil {
		tool.LogERRAdmin(err.Error())
		panic(err.Error())
	}
	_, err = tool.OrmEngine(config)
	if err != nil {
		tool.LogERRAdmin(err.Error())
		return
	}
	app := gin.Default()
	// 页面返回：服务器./packages目录下地文件信息
	// app.Static("/", "./blogpic/")
	// StaticFile 是加载单个文件，而StaticFS 是加载一个完整的目录资源：前一个参数是网络地址的位置，后一个参数是实际位置
	app.StaticFS("/blogimg", http.Dir("E:/finalgo/picturefile/blogpic"))
	app.StaticFS("/portrait", http.Dir("E:/finalgo/picturefile/userpic"))

	Router(app)
	err = app.Run(config.Apphost + ":" + config.Port)
	if err != nil {
		tool.LogERRAdmin("程序启动失败" + err.Error())
	}


}

func Router(route *gin.Engine)  {
	new(controllers.Usercontroller).Router(route)
	new(controllers.Blogcontroller).Router(route)
	new(controllers.Picturecontroller).Router(route)
	new(controllers.Commentcontroller).Router(route)
	new(controllers.Categorycontroller).Router(route)
}
