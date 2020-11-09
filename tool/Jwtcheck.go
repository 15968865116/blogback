package tool
// 用于生成token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	// 声明签名信息
	SigningKey []byte
}

func NewJwt() *JWT {
	return &JWT{
		[]byte("HS256"),
	}
}

//定义错误信息
var (
	TokenExpired     = errors.New("Token 已经过期")
	TokenNotValidYet = errors.New("Token 未激活")
	TokenMalformed   = errors.New("Token 错误")
	TokenInvalid     = errors.New("Token 无效")
)

// 自定义有效载荷(这里采用自定义的Name和Email作为有效载荷的一部分)
type CustomClaims struct {
	Account  string `json:"account"`
	Name string `json:"name"`
	// StandardClaims结构体实现了Claims接口(Valid()函数)
	jwt.StandardClaims
}

// 生成token
func (j *JWT)CreateToken(claim CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	return token.SignedString(j.SigningKey)
}

// 解析token
func (j *JWT)ParseToken(tokenstring string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenstring,&CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token == nil {
		return nil, TokenInvalid
	}
	//解析到Claims 构造中
	if c, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return c, nil
	}
	return nil, TokenInvalid
}

// 生成token并存入redis数据库
func TokenintoRedis(claim CustomClaims) (string,error){
	j := NewJwt()
	token, err := j.CreateToken(claim)
	if err != nil {
		return "", err
	}
	/* 解析token
	c, err := j.ParseToken(token)
	if err != nil{
		return "", err
	}
	 */
	Setjwt(claim.Account,token)

	return token,nil
}
