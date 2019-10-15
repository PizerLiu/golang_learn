package core

import (
	"database/sql"
)

//输出外部要加载的参数
var (
	db   *sql.DB
	es   EsExecute
	etcd *EtcdExecute
)

func Run() {
	//这个可以设置加载哪些core组件
	db = dbInit()
	es = esInit()
	etcd = etcdInit()
}

func exit() {
	//这个可以设置关闭哪些core组件的连接

}

func GetDb() *sql.DB {
	return db
}

func GetEsExecute() EsExecute {
	return es
}

func GetEtcdExecute() *EtcdExecute {
	return etcd
}
