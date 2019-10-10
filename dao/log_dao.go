package dao

import "pizer_project/globle/vo"

type LogDao struct{}

func (m LogDao) SelectLogDao(student vo.Student) int64 {
	return 1
}
