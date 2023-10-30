package repositories

import (
	"fmt"
	"net/http"
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

func GetSavingBy(data *[]models.Saving, id string) error {
	result := configs.DB.Where("member_id = ?", id).Find(&data)
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

func DelSaving(id string, period string) (res interface{}, err error) {
	fmt.Println(id)
	fmt.Println(period)
	result := configs.DB.Unscoped().Debug().Where("member_id = ? and period = ?", id, period).Delete(&models.Saving{})
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return map[string]interface{}{
			"status":  http.StatusNotFound,
			"message": "No savings found with the given member ID and period.",
		}, nil
	}

	return map[string]interface{}{
		"status":       http.StatusOK,
		"message":      "Savings deleted successfully.",
		"rows_deleted": result.RowsAffected,
	}, nil
}
