package routers

import (
	"github.com/gin-gonic/gin"
	"goshiyan/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	return r
}