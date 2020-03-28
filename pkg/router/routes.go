package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// RouteHandler : grouping the routes
type RouteHandler struct {
	Validate *validator.Validate
}

// GetRoutingEngine :
func GetRoutingEngine(version string, rh *RouteHandler) *gin.Engine {
	gin.SetMode(gin.DebugMode)
	ginRouter := gin.Default()
	v := ginRouter.Group(version)
	rh.Service(v)
	return ginRouter
}

// Service :
func (rh *RouteHandler) Service(router *gin.RouterGroup) {
	router.GET("/health", rh.getServerHealth)
	return
}
