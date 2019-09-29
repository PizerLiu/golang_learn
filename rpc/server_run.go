package rpc

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "pizer_project/rpc/proto"
	rpcServer "pizer_project/rpc/server"
)

const port = ":9192"

/**
  1. 首先我们必须实现我们自定义rpc服务,例如:rpc SayHello()-在此我们可以实现我们自己的逻辑
  2. 创建监听listener
  3. 创建grpc的服务
  4. 将我们的服务注册到grpc的server中
  5. 启动grpc服务,将我们自定义的监听信息传递给grpc服务器
*/

//todo rpc注册服务要推给etcd或者zk，client去注册中心拿
//[注意] rpc系统内部调用

func Run() {

	//  创建server端监听端口
	list, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}

	//  创建grpc的server
	server := grpc.NewServer()
	//  注册我们自定义的helloworld服务
	pb.RegisterStudentInfoServer(server, &rpcServer.StudentInfoServer{})

	//  启动grpc服务
	fmt.Println("grpc 服务启动... ...")
	server.Serve(list)

}
