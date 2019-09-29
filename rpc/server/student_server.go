package server

import (
	"context"
	pb "pizer_project/rpc/proto"
	"pizer_project/service"
)

// 定义Service名
var studentService service.StudentService

// 定义struct实现我们自定义的helloworld.proto对应的服务
type StudentInfoServer struct{}

func (StudentInfoServer) GetStudentInfo(ctx context.Context, in *pb.StudentUserRequest) (*pb.StudentUserResponse, error) {
	//去service业务层拿结果，然后用pb返回
	student := studentService.GetStudentById(in.Id)
	return &pb.StudentUserResponse{student.Id, student.Name, student.Ago, student.Sex}, nil
}
