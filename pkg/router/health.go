package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (rh *RouteHandler) getServerHealth(ctx *gin.Context) {
	fmt.Print("Hello world")
	return
}
