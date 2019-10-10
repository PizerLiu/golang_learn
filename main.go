package main

import (
	"fmt"
	"pizer_project/api"
	"pizer_project/conf"
	"pizer_project/core"
	"pizer_project/rpc"
)

// todo
//  1.api日志写入elasticSearch+kibana
//  2.rpc注册服务要推给etcd或者zk，client去注册中心拿；grpc-proxy可以同时访问多个etcd服务器
//  3.优化api加载路由和rpc加载service，现在太机械
//  4.core组件添加【kafka】【redis】
//  5.gev reactor模式tcp网络库
//  6.启动时，扫描pb目录，自动更新pb.go
//  7.将api的handle与rpc使用service统一为一个方法
//  8.main方法后阻塞，api rpc线程同时启动
//  9.配置文件热加载，动态读取更改的配置

// done
//  1.api参数必须校验

func init() {
	//初始化参数配置，dev、online
	fmt.Println("加载配置参数")
	conf.Run()
	//启动core组件
	fmt.Println("加载core组件")
	core.Run()
}

func main() {
	//启动api server服务
	fmt.Println("启动api server服务")
	go api.Run()
	//启动rpc server服务
	//这里可以再做一步：目录中proto文件自动生成pb.go文件
	fmt.Println("启动rpc server服务")
	go rpc.Run()

	select {}
}
