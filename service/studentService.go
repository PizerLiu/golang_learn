package service

import (
	"pizer_project/dao"
	"pizer_project/globle/vo"
)

// 定义dao名
var studentDao dao.StudentDao

// 定义struct实现我们自定义的helloworld.proto对应的服务
type StudentService struct{}

func (m StudentService) AddStudent(student vo.Student) int64 {
	return studentDao.AddStudentDao(student)
}

func (m StudentService) GetStudentById(id int64) vo.Student {
	return studentDao.GetStudentByIdDao(id)
}
