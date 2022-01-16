package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goshiyan/common"
	"goshiyan/routers"
)



func main()  {
	common.InitDB()
	// 1.创建路由
	r := gin.Default()
	routers.CollectRoute(r)
	fmt.Println("http://127.0.0.1:8000")
	r.Run(":8000")
	fmt.Println("work")

}

