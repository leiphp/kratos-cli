/**
 * @package api
 * @file      : greeter.go
 * @author    : LeiXiaoTian
 * @contact   : 1124378213@qq.com
 * @time      : 2023/3/24 16:55
 **/
package api

import (
	"github.com/gin-gonic/gin"
	"testGo/internal/service"
	"testGo/library/common/response"
)

var Greeter = greeterController{}

type greeterController struct{}


func (e *greeterController) SayHello(c *gin.Context) {
	result, err := service.Greeter.SayHello(c)
	response.JsonData(c, result, err)
}

