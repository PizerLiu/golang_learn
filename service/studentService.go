package service

import (
	"pizer_project/core"
	"pizer_project/dao"
	"pizer_project/globle/vo"
)

// 定义dao名
var studentDao dao.StudentDao

type StudentService struct{}

func (m StudentService) AddStudent(student vo.Student) int64 {
	//es中添加数据
	es := core.GetEsExecute()
	es.AddSourceByIndex("pizer", `{"age":"20"}`)

	return studentDao.AddStudentDao(student)
}

func (m StudentService) GetStudentById(id int64) vo.Student {
	return studentDao.GetStudentByIdDao(id)
}
