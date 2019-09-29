package controller

import (
	"github.com/gin-gonic/gin"
	apiUtil "pizer_project/api/utils"
	"pizer_project/globle/vo"
	"pizer_project/service"
)

/************************************
*	message: 记录student业务相关controller方法
*	author: pizer
*	time: 2019/9/29
************************************/

// 定义Service名
var studentService service.StudentService

type StudentController struct{}

// 新增student数据
func (StudentController) AddStudentHandle(c *gin.Context) {
	var studentInfo vo.Student
	if apiUtil.CheckParam(c, &studentInfo) {
		//检查参数有问题
		return
	}
	studentService.AddStudent(studentInfo)
	apiUtil.ApiSucc(c, "添加同学成功")
	return
}
