package database

import (
	models2 "periferia_it_social_network/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	var count int64
	DB.Model(&models2.User{}).Count(&count)
	if count > 0 {
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	user := models2.User{
		Name:      "John",
		LastName:  "Perez",
		Username:  "JohnP",
		Email:     "juan@test.com",
		Password:  string(password),
		BirthDate: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
	}

	DB.Create(&user)

	post := models2.Post{
		Message: "Create a new post",
		UserID:  user.ID,
	}

	DB.Create(&post)
}
