package main

import (
	"github.com/gin-gonic/gin"
	"will/common"
	"will/router"
)

func main() {
	common.InitDB()
	r := gin.Default()

	r = router.WillCollectRoute(r)

	panic(r.Run(":9000"))
}
