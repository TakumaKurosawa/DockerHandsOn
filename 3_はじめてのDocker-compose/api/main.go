package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{"*"},
	}))

	r.GET("/browser/access", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "access from web browser.",
		})
	})
	r.GET("/server/access", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "access from server.",
		})
	})
	r.Run()
}
