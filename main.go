package main

import (
	"finalgo/controllers"
	"finalgo/tool"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)


func main()  {

	config, err := tool.ParseConfig("./config/config.json")
	if err != nil {
		panic(err.Error())
	}
	_, err = tool.OrmEngine(config)
	if err != nil {
		log.Error().Err(err)
	}
	app := gin.Default()
	Router(app)
	app.Run(config.Apphost + ":" + config.Port)


}

func Router(route *gin.Engine)  {
	new(controllers.Usercontroller).Router(route)
	new(controllers.Blogcontroller).Router(route)
}
