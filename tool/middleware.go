package tool

import (
	"github.com/gin-gonic/gin"
)

func Tokencheck(c *gin.Context) {
	token := c.GetHeader("token")
	account := c.GetHeader("puberaccount")

	result, err := Getjwt(account, token)
	if err != nil {
		c.AbortWithStatusJSON(400, map[string]interface{}{
			"msg": "出现未知错误",
		})
	}
	if result == true {
		c.Next()
	} else {
		c.AbortWithStatusJSON(400, map[string]interface{}{
			"msg": "验证失败",
		})
	}

}
