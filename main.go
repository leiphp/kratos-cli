package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	v1 "testGo/grpc/helloworld/v1"
	"testGo/internal/api"
)

func main() {
	// grpc服务
	var serviceHost = "127.0.0.1:9000"
	conn, err := grpc.Dial(serviceHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	client := v1.NewGreeterClient(conn)
	greeterApi := api.NewGreeterController(client)
	defer conn.Close()


	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	r.GET("/check/health", health)
	// 路由组1 ，处理GET请求
	v1 := r.Group("/v1")
	// {} 是书写规范
	{
		v1.GET("/greeter/say_hello", api.Greeter.SayHello)
		v1.GET("/greeter/GrpcSayHello",  greeterApi.GrpcSayHello)
	}
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8001")
}

//健康检查
func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": []int{},
		"msg": "service is ok",
	})
}