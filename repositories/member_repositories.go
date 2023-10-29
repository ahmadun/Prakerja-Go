package repositories

import (
	"prakerja12/configs"
	"prakerja12/models"
)

func GetMembers(data *[]models.Member) error {
	result := configs.DB.Find(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddMember(data *models.Member) error {
	result := configs.DB.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
