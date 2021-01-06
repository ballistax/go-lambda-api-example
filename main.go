package main

import "github.com/gin-gonic/gin"

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	app := gin.Default()

	app.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	app.POST("/user", func(c *gin.Context) {
		var u user
		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		c.JSON(200, u)
	})

	app.Run(":3000")
}
