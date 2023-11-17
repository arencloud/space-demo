package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    int
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(db *gorm.DB, user *User) (result *gorm.DB) {
	result = db.Create(user)

	return result
}

func GetUsers(db *gorm.DB, user *[]User) (result *gorm.DB) {
	result = db.Find(user)
	return
}

func GetUser(db *gorm.DB, user *User, id int) (result *gorm.DB) {
	result = db.Where("id = ?", id).First(user)
	return result
}

func UpdateUser(db *gorm.DB, user *User) (result *gorm.DB) {
	result = db.Save(user)
	return result
}

func DeleteUser(db *gorm.DB, user *User, id int) (result *gorm.DB) {
	result = db.Where("id = ?", id).Delete(user)

	return result
}
