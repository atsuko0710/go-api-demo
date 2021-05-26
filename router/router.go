package router

import (
	"go-api-demo/handler/sd"
	"go-api-demo/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// 在处理某些请求时可能因为程序bug或者其他异常情况导致程序panic，这时候为了不影响下⼀次请求的调⽤，需要通过gin.Recovery()来恢复API服务器
	g.Use(gin.Recovery())

	// 添加中间件
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)

	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// 健康检查
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		// svcd.GET("/disk", sd.DiskCheck)
		// svcd.GET("/cpu", sd.CPUCheck)
		// svcd.GET("/ram", sd.RAMCheck)
	}
	return g
}
