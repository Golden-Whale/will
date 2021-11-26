package router

import (
	"github.com/gin-gonic/gin"
	"will/controller"
	"will/middleware"
)

func WillCollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.GET("/api/will", controller.GetWill)
	r.PUT("/api/will", controller.PutWill)

	r.PUT("/api/will/date/range", controller.GetWillHourRange)
	return r
}
