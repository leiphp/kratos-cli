package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main22() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}

//git flow feature finish myfeature

//git flow release


var logger *log.Logger


func init() {
	//file, err := os.Create("testlog.log")

	//if err != nil {
	//	log.Fatal(err)
	//}
	// log.LstdFlags 设置初始值：相当于 log.Ldate|log.Ltime
	//log.Llongfile 显示完整的文件名和行数 除了这之外还有 Lmicroseconds
	// Llongfile Lshortfile LUTC
	// log.new 有三个参数，第一个输出位置，第二个为日志输出前缀，第三个设置logger的属性
	logger = log.New(os.Stdout, "[Info]", log.LstdFlags|log.Llongfile)
}


func main() {
	logger.Println("begin TestLog ...")
	test()
	log.SetFlags(log.LstdFlags | log.Llongfile)
	// 输出到标准输出
	log.Println("1.Println log with log.LstdFlags1 ...")
	// 输出到定义的输出位置
	logger.Println("2.Println log with log.LstdFlags ...")
	//logger.SetFlags(log.LstdFlags)
	test()


}

func test() {
	logger.Println("helloworld")
}
