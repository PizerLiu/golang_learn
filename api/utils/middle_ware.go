package utils

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"pizer_project/conf"
	"pizer_project/globle"
)

//jwt鉴权
//该中间件打开后
//[注意] header中入 Authorization = token
func Jwt() gin.HandlerFunc {
	//加载时的逻辑
	apiConf := conf.GetBaseConf()
	apiSecretKey := apiConf.Get(globle.CONST_CONFIG_SECTION_API, "apiSecretKey")
	//中间件逻辑写在Auth中
	return func(c *gin.Context) {
		//根据token获取userId，并储存到变量中。c.Set("operatorId", 1111)即可

		_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := []byte(apiSecretKey)
			return b, nil
		})

		if err != nil {
			c.AbortWithError(401, err)
		}
	}
}
