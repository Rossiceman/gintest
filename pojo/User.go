package pojo

import (
	"golangApi/database"
)

// type User struct {
// 	Id       int    `json:"UserId"`
// 	Name     string `json:"UserName"`
// 	Password string `json:"UserPassword"`
// 	Email    string `json:"UserEmail"`
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
func DeleteUser(userId string) User {
	user := User{}
	database.DBConnect.Where("id=?", userId).Delete(&user)
	return user
}

// UpdataUser
func UpdataUser(userId string, user User) User {
	database.DBConnect.Model(&user).Where("id=?", userId).Updates(user)
	return user
}
