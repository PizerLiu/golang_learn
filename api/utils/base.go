package utils

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"pizer_project/globle"
)

func init() {
	//开启对struct json参数校验
	//govalidator.SetFieldsRequiredByDefault(true)
}

// check params
// example : checkParam(c, &studentInfo)
func CheckParam(c *gin.Context, param interface{}) bool {
	err := c.BindJSON(&param)

	if err != nil {
		ApiFatal(c, globle.ENUM_PARAM_ERROR, param)
		return true
	}

	//自定义tag验证函数
	if _, err := govalidator.ValidateStruct(param); err != nil {
		ApiFatal(c, globle.ENUM_CHECK_PARAM_ERROR, err)
		return true
	}

	return false
}

// 捕捉 panic
func RecoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(500, gin.H{"code": -1, "message": err})
}

func NoRouter(c *gin.Context) {
	//c.HTML(http.StatusNotFound, "error.html", "Sorry,I lost myself!")
	c.JSON(404, gin.H{"code": -1, "message": "Not router"})
}

func NoMethod(c *gin.Context) {
	c.JSON(404, gin.H{"code": -1, "message": "Not method"})
}

func ApiSucc(c *gin.Context, info interface{}) {
	c.JSON(200, gin.H{"code": 0, "message": info})
}

func ApiFatal(c *gin.Context, errorCode int, info interface{}) {
	c.JSON(200, gin.H{"code": errorCode, "message": info})
}
