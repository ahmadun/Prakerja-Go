package repositories

import (
	"prakerja12/configs"
	"prakerja12/models"
)

func GetLoanPaymentBy(data *[]models.LoanPayment, id string) error {
	result := configs.DB.Where("member_id = ?", id).Find(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
