package dao

import (
	"fmt"
	"log"
	"pizer_project/core"
	"pizer_project/globle/vo"
)

// todo
//  1.数据库回滚、事务加上
//  2.insert、update、delete成功返回true，失败返回false

type StudentDao struct{}

func (m StudentDao) AddStudentDao(student vo.Student) int64 {
	db := core.GetDb()
	//事务
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()

	stmt1, err := tx.Prepare("insert student set name =?,ago=?,sex=?")
	defer stmt1.Close()
	if err != nil {
		log.Fatal(err)
	}

	result, err := stmt1.Exec(&student.Name, &student.Ago, &student.Sex)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func (m StudentDao) GetStudentByIdDao(id int64) vo.Student {
	db := core.GetDb()
	var student vo.Student

	rows, err := db.Query("select * from student where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	//返回字段
	for rows.Next() {
		err = rows.Scan(
			&student.Id,
			&student.Name,
			&student.Sex,
			&student.Ago,
			&student.CreateTime)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("student=", student)
	}

	//return result
	return student
}
