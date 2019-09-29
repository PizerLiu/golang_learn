package core

import (
	"database/sql"
)

//输出外部要加载的参数
var (
	db  *sql.DB
	exe EsExecute
)

func Run() {
	//这个可以设置加载哪些core组件
	db = dbInit()
	//exe = esInit()
}

func exit() {
	//这个可以设置关闭哪些core组件的连接

}

func GetDb() *sql.DB {
	return db
}

func GetEsExecute() EsExecute {
	return exe
}
