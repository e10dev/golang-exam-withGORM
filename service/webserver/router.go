package webserver

import (
	"e10dev.example/service/webserver/api"
	"e10dev.example/service/webserver/api/account"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())
	a := r.Group("/api")
	{
		a.POST("/login", api.Login)
	}
	b := r.Group("/api/account")
	{
		b.GET("", account.GetAllAccount)
		b.GET("/:seq", account.GetAccount)
		b.POST("", account.NewAccount)
		b.PUT("/:seq", account.UpdateAccount)
		b.DELETE("/:seq", account.DeleteAccount)
		b.GET("/download", account.DownloadCSV)
	}
	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
