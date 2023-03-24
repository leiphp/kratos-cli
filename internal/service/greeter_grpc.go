/**
 * @package service
 * @file      : greeter_grpc.go
 * @author    : LeiXiaoTian
 * @contact   : 1124378213@qq.com
 * @time      : 2023/3/24 18:14
 **/
package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "testGo/grpc/helloworld/v1"
)


var GreeterGrpc = greeterGrpcService{}

type greeterGrpcService struct{}

// GRPC调用
func (s *greeterGrpcService) GrpcSayHello(c *gin.Context, client v1.GreeterClient) (interface{}, error) {
	rsp, err := client.SayHello(c, &v1.HelloRequest{
		Name: "LeiPHP",
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return rsp, nil
}