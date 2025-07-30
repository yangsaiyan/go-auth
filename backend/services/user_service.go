package services

import (
	"errors"
	"fmt"
	"go-auth/database"
	"go-auth/models"
	"go-auth/utils"
)

func Login(email string, password string) (models.User, error) {
	fmt.Println(email, password)
	var user models.User

	if database.DB.Where("email = ?", email).First(&user).Error != nil {
		return models.User{}, errors.New("user not found")
	}

	if !utils.VerifyPassword(password, user.Password) {
		return models.User{}, errors.New("invalid password")
	}

	_, err := utils.GenerateJWTCookie(user.ID)
	if err != nil {
		return models.User{}, errors.New("error generating cookie")
	}

	return user, nil
}

func Signup(email string, password string) (models.User, error) {
	var err error

	user := models.User{
		Email:    email,
		Password: password,
	}

	if database.DB.Where("email = ?", email).First(&user).Error != nil {
		return models.User{}, errors.New("user already exists")
	}

	user.Password, err = utils.HashPassword(password)
	if err != nil {
		return models.User{}, errors.New("error hashing password")
	}

	_, err = utils.GenerateJWTCookie(user.ID)
	if err != nil {
		return models.User{}, errors.New("error generating cookie")
	}

	database.DB.Create(&user)

	return user, nil
}

func Update(email string, password string) (models.User, error) {
	user := models.User{
		Email:    email,
		Password: password,
	}

	if email == "" {
		return models.User{}, errors.New("email is required")
	}

	if password == "" {
		return models.User{}, errors.New("password is required")
	}

	err := database.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	if !utils.VerifyPassword(password, user.Password) {
		return models.User{}, errors.New("invalid password")
	}

	user.Password, err = utils.HashPassword(password)
	if err != nil {
		return models.User{}, errors.New("error hashing password")
	}

	database.DB.Save(&user)

	return user, nil
}

func Delete(email string) (models.User, error) {
	var user models.User

	if database.DB.Where("email = ?", email).First(&user).Error != nil {
		return models.User{}, errors.New("user not found")
	}

	database.DB.Delete(&user)

	return user, nil
}

func GetUserByID(userID uint) (models.User, error) {
	var user models.User

	if database.DB.Where("id = ?", userID).First(&user).Error != nil {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}
