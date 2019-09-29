package routers

import (
	"github.com/DeanThompson/ginpprof"
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	"pizer_project/api/controller"
	apiUtil "pizer_project/api/utils"
	"pizer_project/conf"
	"pizer_project/globle"
	"pizer_project/utils"
)

var apiConf *utils.CfgFileConfig

// 定义Controller名
var baseController controller.BaseController
var studentController controller.StudentController

func InitRun() {
	//读取配置信息
	apiConf := conf.GetBaseConf()
	apiSecretKey := apiConf.Get(globle.CONST_CONFIG_SECTION_API, "apiSecretKey")
	apiTls := apiConf.Get(globle.CONST_CONFIG_SECTION_API, "apiTls")
	apiTlsAddr := apiConf.Get(globle.CONST_CONFIG_SECTION_API, "apiTlsAddr")
	apiAddr := apiConf.Get(globle.CONST_CONFIG_SECTION_API, "apiAddr")

	engine := gin.Default()
	//程序监控 http://127.0.0.1:8080/debug/pprof/
	ginpprof.Wrap(engine)

	engine.Use(nice.Recovery(apiUtil.RecoveryHandler))

	// 未知路由处理
	engine.NoRoute(apiUtil.NoRouter)
	// 未知调用方式
	engine.NoRoute(apiUtil.NoMethod)

	// 公开路由
	public := engine.Group("/api/v1")
	{
		public.POST("/auth", baseController.CreateTokenHandle)
	}

	// 非公开路由
	private := engine.Group("/api/v1/private")
	{
		//https://studygolang.com/articles/10628?fr=sidebar
		//[注意] header中入 Authorization = token
		private.Use(jwt.Auth(apiSecretKey))
		private.POST("/add_student", studentController.AddStudentHandle)
	}

	//service run
	if apiTls == "true" {
		engine.RunTLS(apiTlsAddr, "keys/server.crt", "keys/server.key")
	} else {
		engine.Run(apiAddr)
	}

}
