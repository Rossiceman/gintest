package src

import (
	"golangApi/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")

	//POST與GET的方法，在postman中的POST輸入值，再從GET中取出值
	user.GET("/", service.FindAllUsers)
	user.GET("/:id", service.FindAllUsers)
	user.POST("/", service.PostUser)

	//delete user

	user.DELETE("/:id", service.DeleteUser)

	//put user

	user.PUT("/:id", service.PutUser)
}
