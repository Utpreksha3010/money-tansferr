package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	r.GET("/akansha", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.GET("/tik", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "tok",
		})
	})

	r.GET("/shiv", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "akansha",
		})
	})
	r.Run()

}
