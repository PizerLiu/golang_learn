package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"pizer_project/conf"
	"pizer_project/core"

	pb "pizer_project/rpc/proto"
)

// 此处应与服务器端对应
const address = "127.0.0.1:9192"

/**
  1. 创建groc连接器
  2. 创建grpc客户端,并将连接器赋值给客户端
  3. 向grpc服务端发起请求
  4. 获取grpc服务端返回的结果
*/

func init() {
	//初始化参数配置，dev、online
	fmt.Println("加载配置参数")
	conf.Run()
	//启动core组件
	fmt.Println("加载core组件")
	core.Run()
}

func main() {
	etcds := core.GetEtcdExecute()
	sre := etcds.GetServerRegisterByName("StudentInfoServer")
	fmt.Println(sre)

	// 创建一个grpc连接器
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	// 当请求完毕后记得关闭连接,否则大量连接会占用资源
	defer conn.Close()

	// 创建grpc客户端
	c := pb.NewStudentInfoClient(conn)

	// 客户端向grpc服务端发起请求
	result, err := c.GetStudentInfo(
		context.Background(),
		&pb.StudentUserRequest{
			Id: 40,
		})
	if err != nil {
		fmt.Println("请求失败!!!", err)
		return
	}
	// 获取服务端返回的结果
	fmt.Println(result.Name)
}
