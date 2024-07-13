package main

import (
	"golangApi/database"
	"golangApi/src"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	v1 := router.Group("/v1")
	src.AddUserRouter(v1)

	go func() {
		database.DD()
	}()

	router.Run(":80")
}

// //router := gin.Default()
// //router.GET("/ping", func(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message":  "TEST",
// 		"message1": "1123",
// 	})
// })
// router.POST("/ping/:id", func(c *gin.Context) {
// 	id := c.Param("id")
// 	c.JSON(200, gin.H{
// 		"id": id,
// 	})
// })*/
