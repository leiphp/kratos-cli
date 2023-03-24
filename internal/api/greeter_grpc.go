/**
 * @package api
 * @file      : greeter_grpc.go
 * @author    : LeiXiaoTian
 * @contact   : 1124378213@qq.com
 * @time      : 2023/3/24 18:17
 **/
package api

import (
	"github.com/gin-gonic/gin"
	v1 "testGo/grpc/helloworld/v1"
	"testGo/internal/service"
	"testGo/library/common/response"
)

type GreeteController struct {
	client v1.GreeterClient
}

func NewGreeterController(client v1.GreeterClient) *GreeteController {
	return &GreeteController{client: client}
}

func (e *GreeteController) GrpcSayHello(c *gin.Context) {
	result, err := service.GreeterGrpc.GrpcSayHello(c, e.client)
	response.JsonData(c, result, err)
}
