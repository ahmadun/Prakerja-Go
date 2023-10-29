package repositories

import (
	"errors"
	"prakerja12/configs"
	"prakerja12/models"

	"golang.org/x/crypto/bcrypt"
)

func Register(user *models.User) error {
	result := configs.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func Login(user *models.UserLogin) (models.User, error) {

	var users *models.User

	result := configs.DB.Debug().Model(models.User{}).Where("email = ?", user.Email).First(&users)
	if result.Error != nil {
		return *users, result.Error
	}

	if result.RowsAffected == 0 {
		return *users, errors.New("no user found with the given email")
	}

	err := models.VerifyPassword(users.Password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return *users, errors.New("invalid password")
	}

	return *users, nil
}
