package main

import (
	"errors"
	"fmt"
	"gorm/database"
	"gorm/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()
	// CreateUser("fajri010217@gmail.com")
	// getUserById(1)
	// updateUser(1, "emailupdated@gmail.com")
	// createProduct(1, "ylo", "llLL")
	// getUsersWithProducts()
	deleteProductById(2)
}

func CreateUser(email string) {
	db := database.GetDb()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error

	if err != nil {
		fmt.Println("Error creating user data:", err.Error())
		return
	}

	fmt.Println("New User Data:", User)
}

func getUserById(id int) {
	db := database.GetDb()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}

		print("Error finding user:", err)
	}

	fmt.Printf("User data: %+v \n", user)
}

func updateUser(id uint, email string) {
	db := database.GetDb()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(
		models.User{Email: email}).Error

	if err != nil {
		fmt.Println("Error updating user data:", err)
	}

	fmt.Printf("Update user's email: %+v \n", user.Email)
}

func createProduct(userId uint, brand string, name string) {
	db := database.GetDb()
	product := models.Product{
		Name:   name,
		Brand:  brand,
		UserID: userId,
	}

	err := db.Create(&product).Error

	if err != nil {
		fmt.Println("Error creating product data:", err.Error())

		return
	}

	fmt.Println("New Product Data:", product)
}

func getUsersWithProducts() {
	db := database.GetDb()

	users := models.User{}

	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting datas with products:", err.Error())
		return
	}

	fmt.Println("User Datas with Products")
	fmt.Printf("%+v", users)
}

func deleteProductById(id uint) {
	db := database.GetDb()

	product := models.Product{}

	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deletin gproduct", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been successfully deleted", id)
}
