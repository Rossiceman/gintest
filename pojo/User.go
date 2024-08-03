package pojo

import (
	"golangApi/database"
	"log"
)

// type User struct {
// 	Id       int    `json:"UserId" binding:"required"`
// 	Name     string `json:"UserName" binding:"required,gt=5"`
// 	Password string `json:"UserPassword" binding:"min=4,max=20,userpasd"`
// 	Email    string `json:"UserEmail" binding:"email"`
// }

type User struct {
	Id       int
	Name     string
	Password string
	Email    string
}

func FindAllUsers() []User {
	var user []User
	database.DBConnect.Find(&user)
	return user
}

func FindByUserId(userId int) User {
	var user User
	database.DBConnect.Where("id=?", userId).First(&user)
	return user
}

// Post
func CreateUser(user User) User {
	database.DBConnect.Create(&user)
	return user
}

// DeleteUser
func DeleteUser(userId string) bool {
	user := User{}
	result := database.DBConnect.Where("id=?", userId).Delete(&user)
	log.Println(result)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

// UpdataUser
func UpdataUser(userId string, user User) User {
	database.DBConnect.Model(&user).Where("id=?", userId).Updates(user)
	return user
}
