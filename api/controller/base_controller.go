package controller

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	apiUtil "pizer_project/api/utils"
	"pizer_project/conf"
	"pizer_project/globle"
	"time"
)

/************************************
*	message: 记录基础controller方法
*	author: pizer
*	time: 2019/9/29
************************************/

// 定义router名
//var baseRouter routers.BaseRouter

//注意：
//	1.方法名字首字母为小写时，不可被其他非本目录下文件调用
type BaseController struct{}

// 生成 token
func (m BaseController) CreateTokenHandle(c *gin.Context) {
	//读取配置信息
	apiConf := conf.GetBaseConf()
	apiUsername := apiConf.Get(globle.CONST_CONFIG_SECTION_API, "apiUsername")
	apiPassword := apiConf.Get(globle.CONST_CONFIG_SECTION_API, "apiPassword")
	apiSecretKey := apiConf.Get(globle.CONST_CONFIG_SECTION_API, "apiSecretKey")

	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == apiUsername && password == apiPassword {
		// Create the token
		token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
		// Set some claims
		token.Claims = jwt_lib.MapClaims{
			"Id":  "Christopher",
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		}
		// Sign and get the complete encoded token as a string
		tokenString, err := token.SignedString([]byte(apiSecretKey))
		if err != nil {
			apiUtil.ApiFatal(c, globle.ENUM_HEADER_TOKEN_ERROR, "Could not generate token")
			return
		}
		apiUtil.ApiSucc(c, gin.H{"token": tokenString})
		return
	} else {
		apiUtil.ApiFatal(c, globle.ENUM_CHECK_PARAM_ERROR, "username or password is wrong")
		return
	}

}
