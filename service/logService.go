package service

import "pizer_project/dao"

// 定义dao名
var logDao dao.LogDao

// 定义struct实现我们自定义的helloworld.proto对应的服务
type LogService struct{}
