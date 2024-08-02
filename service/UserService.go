package service

import (
	"golangApi/pojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var userList = []pojo.User{}

//Get User

func FindAllUsers(c *gin.Context) {
	data := pojo.FindAllUsers()
	c.JSON(http.StatusOK, data)
}

//Post User

func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error : "+err.Error())
		return
	}
	newUser := pojo.CreateUser(user)
	c.JSON(http.StatusOK, newUser)

}

//delete User

func DeleteUser(c *gin.Context) {
	// userId, _ := strconv.Atoi(c.Param("id"))
	// for _, user := range userList {
	// 	log.Println(user)

	// 	userList = append(userList[:userId], userList[userId+1:]...)
	// 	c.JSON(http.StatusOK, "Successfully deleted")
	// 	return

	// }
	// c.JSON(http.StatusNotFound, "Error")

	user := pojo.DeleteUser(c.Param("id"))
	if !user {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, "Successfully")
}

//Update User

func PutUser(c *gin.Context) {
	// beforeUser := pojo.User{}
	// err := c.BindJSON(&beforeUser)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, "Error")
	// }
	// userId, _ := strconv.Atoi(c.Param("id"))
	// for key, user := range userList {
	// 	if userId == user.Id {
	// 		userList[key] = beforeUser
	// 		log.Println(userList[key])
	// 		c.JSON(http.StatusOK, "Successfully")
	// 		return
	// 	}
	// }
	// c.JSON(http.StatusNotFound, "Error")

	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}
	user = pojo.UpdataUser(c.Param("id"), user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, user)
}
