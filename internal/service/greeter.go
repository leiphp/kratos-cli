/**
 * @package service
 * @file      : greeter.go
 * @author    : LeiXiaoTian
 * @contact   : 1124378213@qq.com
 * @time      : 2023/3/24 16:56
 **/
package service

import (
	"github.com/gin-gonic/gin"
)

var Greeter = greeterService{}

type greeterService struct{}


// 打招呼
func (s *greeterService) SayHello(c *gin.Context) (interface{}, error) {
	name := c.DefaultQuery("name", "jack")
	return name, nil
}
