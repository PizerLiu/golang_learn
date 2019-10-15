package routers

import (
	"fmt"
	"github.com/DeanThompson/ginpprof"
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	"os"
	"pizer_project/api/controller"
	apiUtil "pizer_project/api/utils"
	"pizer_project/conf"
	"pizer_project/globle"
	"pizer_project/utils"
	"strings"
)

var apiConf *utils.CfgFileConfig

// 定义Controller名
var baseController controller.BaseController
var studentController controller.StudentController

var file *os.File

func InitRun() {
	//读取配置信息
	apiConf := conf.GetBaseConf()
	apiTls := apiConf.Get(globle.CONST_CONFIG_SECTION_API, "apiTls")
	apiTlsAddr := apiConf.Get(globle.CONST_CONFIG_SECTION_API, "apiTlsAddr")
	apiAddr := apiConf.Get(globle.CONST_CONFIG_SECTION_API, "apiAddr")

	// gin.Default()默认加载Logger和Recovery中间件
	// Logger中间件: 打印接口请求logger，打印指定样式的log
	// [例子]：[GIN] 2019/10/09 - 14:54:15 | 200 |  243.925229ms |  127.0.0.1 | POST  /api/v1/private/add_student
	// Recovery中间件
	// Recovery 中间件从任何 panic 恢复，如果出现 panic，它会写一个 500 错误
	engine := gin.New()
	//程序监控 http://127.0.0.1:8080/debug/pprof/
	ginpprof.Wrap(engine)

	// api log日志落盘，filebeat再将日志同步至es中
	// 日志切分是放在文件处切分，还是同步至es时再切分？
	// 目的：每条HTTP请求日志，都对应一次用户的请求行为，记录每一条用户请求日志，
	// 对于我们追踪用户行为，过滤用户非法请求，排查程序运行产生的各种问题至关重要，
	// 因此，开发Web应用时一定要记录用户请求行为，并且定时分析过滤。
	logFile := utils.CurrentFile() + "/../log/api_access.log"
	exist, _ := utils.PathExists(logFile)
	if !exist {
		_ = os.Mkdir(utils.CurrentFile()+"/../log", os.ModePerm)
		file, _ = os.Create(logFile)
	} else {
		file, _ = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	}

	c := gin.LoggerConfig{
		//Output: io.MultiWriter(os.Stdout, file)
		Output: file,
		//SkipPaths:[]string{"/api"},
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - [%s] %s %s %s %d %s %s ; [error]=^%s^ \n",
				params.ClientIP,
				params.TimeStamp.Format("2006/01/02 15:04:05"),
				params.Method,
				params.Path,
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
				params.Request.UserAgent(),
				strings.Replace(params.ErrorMessage, "\n", "", -1),
			)
		},
	}

	engine.Use(gin.LoggerWithConfig(c))
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
	private := engine.Group("/api/v1/private", apiUtil.Jwt())
	{
		//https://studygolang.com/articles/10628?fr=sidebar
		//[注意] header中入 Authorization = token
		//private.Use(jwt.Auth(apiSecretKey))
		private.POST("/add_student", studentController.AddStudentHandle)
	}

	//service run
	fmt.Println("api 服务启动...")
	if apiTls == "true" {
		engine.RunTLS(apiTlsAddr, "keys/server.crt", "keys/server.key")
	} else {
		engine.Run(apiAddr)
	}

}
