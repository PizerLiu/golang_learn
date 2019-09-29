package core

import (
	"bytes"
	"database/sql"
	"fmt"
	"pizer_project/conf"
	"pizer_project/globle"
	"pizer_project/utils"
)

import (
	_ "github.com/go-sql-driver/mysql" // 注册mysql driver
)

var dbConf *utils.CfgFileConfig

// 初始化DB
// todo
//  1.数据库连接可以改为连接池
func dbInit() *sql.DB {
	var DB *sql.DB

	dbConf = conf.GetBaseConf()

	hostname := dbConf.Get(globle.CONST_CONFIG_SECTION_DB, "hostname")
	port := dbConf.Get(globle.CONST_CONFIG_SECTION_DB, "port")
	dbname := dbConf.Get(globle.CONST_CONFIG_SECTION_DB, "dbname")
	username := dbConf.Get(globle.CONST_CONFIG_SECTION_DB, "username")
	password := dbConf.Get(globle.CONST_CONFIG_SECTION_DB, "password")

	DB, _ = DbConnect(hostname, port, dbname, username, password)

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	return DB
}

// 连接mysql数据库
func DbConnect(host string, port string, dbname string, username string, password string) (*sql.DB, error) {

	var buf bytes.Buffer
	buf.WriteString(username)
	buf.WriteString(":")
	buf.WriteString(password)
	buf.WriteString("@tcp(")
	buf.WriteString(host)
	buf.WriteString(":")
	buf.WriteString(port)
	buf.WriteString(")/")
	buf.WriteString(dbname)
	buf.WriteString("?charset=utf8&parseTime=true")

	connection := buf.String()

	db, err := sql.Open("mysql", connection)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
