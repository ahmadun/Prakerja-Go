package repositories

import (
	"prakerja12/configs"
	"prakerja12/models"
)

func GetSavings(data *[]models.Saving) error {
	result := configs.DB.Find(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddSaving(data *models.Saving) error {
	result := configs.DB.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
