package controllers
// 控制器，用于提供路径以及相应路径的处理
import (
	"finalgo/dao"
	"finalgo/model"
	"finalgo/tool"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"time"
)

type Usercontroller struct {

}

//实现查询所需的结构
type Usermessagestruct struct {
	Account string `json:"account"`
	Token string `json:"token"`
}

func (uc *Usercontroller) Router(engine *gin.Engine)  {
	engine.POST("/user/Insertuser", uc.Insertuser)
	engine.POST("/user/Login", uc.Selectuser)
	engine.POST("/user/getinfo", uc.Usermessage)
}

//创建新用户
func (uc *Usercontroller) Insertuser(context *gin.Context)  {
	var user model.User
	err := context.BindJSON(&user)
	if err!=nil {
		log.Error().Err(err)
	}
	userdao := dao.Userdao{tool.DBengine}
	result := userdao.Insertuser(user)
	if result > 0 {
		context.JSON(200, map[string]interface{}{
			"code":1,
			"msg":"插入数据成功",
		})
	} else {
		context.JSON(300,map[string]interface{} {
			"code":0,
			"msg":"插入数据失败",
		})
	}
}

//查询用户 即实现登录
func  (uc *Usercontroller)Selectuser(context *gin.Context)  {
	var user model.User
	err := context.BindJSON(&user)
	if err!=nil {
		log.Error().Err(err)
	}
	userdao := dao.Userdao{tool.DBengine}
	usertwo := userdao.Selectuser(user.Account,user.Password)
	if usertwo == nil {
		context.JSON(250,map[string]interface{}{
			"code":0,
			"msg":"登录失败",
			"token":"none",
		})
	} else {
		claim := tool.CustomClaims{
			Account: usertwo.Account,
			Name: usertwo.Name,
			StandardClaims: jwt.StandardClaims{
				NotBefore: time.Now().Unix() - 1000,       //签名生效时间
				ExpiresAt: time.Now().Unix() + 60*60*24*7, //签名过期时间
				Issuer:    "zjj",                          //签名颁发者
			},
		}
		token, err := tool.TokenintoRedis(claim)
		if err != nil {
			context.JSON(200,map[string]interface{} {
				"code":3,
				"msg":"redis缓存失败",
				"token":"None",
			})
		} else{
			context.JSON(200,map[string]interface{} {
				"code":1,
				"username":usertwo.Name,
				"msg":"登录成功",
				"token":token,
			})

		}
	}
}

// 删除用户

// 更新用户资料

// 查询用户资料
func (uc *Usercontroller)Usermessage(context *gin.Context) {
	var usermessage Usermessagestruct
	err := context.BindJSON(&usermessage)
	if err != nil {
		log.Err(err)
	}
	// 判断token情况
	tokenture, err := tool.Getjwt(usermessage.Account,usermessage.Token)
	if err != nil {
		return
	}
	if tokenture != true {
		context.JSON(250,map[string]interface{}{
			"code": 0,
			"msg":"token错误或者登录过期",
		})
	} else{
		userdao := dao.Userdao{tool.DBengine}
		usertwo := userdao.SelectuserMessage(usermessage.Account)
		context.JSON(200,map[string]interface{}{
			"code": 1,
			"msg":"获取成功",
			"name":usertwo.Name,
		})
	}
}