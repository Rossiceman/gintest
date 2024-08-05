package main

import (
	"golangApi/database"
	"golangApi/middlewares"
	"golangApi/pojo"
	"golangApi/src"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func setupLogging() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	setupLogging()

	router := gin.Default()

	// 註冊 Validator func
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("userpasd", middlewares.UserPasd)
		v.RegisterStructValidation(middlewares.UserList, pojo.Users{})
	}

	router.Use(gin.Recovery(), middlewares.Logger())

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
