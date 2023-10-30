package repositories

import (
	"fmt"
	"net/http"
	"prakerja12/configs"
	"prakerja12/models"
	"time"
)

func GetLoans(data *[]models.Loan) error {
	result := configs.DB.Find(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetLoanBy(data *[]models.Loan, id string) error {
	result := configs.DB.Where("member_id = ?", id).Find(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ZeroTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func AddLoan(data *models.Loan) error {
	data.LoanDate = ZeroTime(time.Now())
	result := configs.DB.Create(&data)

	currentDate := data.CreditStart
	paytime := 1
	for currentDate.Before(data.CreditFinish) || currentDate.Equal(data.CreditFinish) {
		payment := models.LoanPayment{
			MemberID:    data.MemberID,
			LoanDate:    data.LoanDate,
			Period:      fmt.Sprintf("%d%02d", currentDate.Year(), currentDate.Month()),
			PaymentTime: paytime,
			Amount:      data.CreditMonthly,
			Remark:      "",
			EntryBy:     data.EntryBy,
		}

		if err := configs.DB.Create(&payment).Error; err != nil {
			return err
		}
		paytime++
		currentDate = currentDate.AddDate(0, 1, 0)
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DelLoan(id string, period string) (res interface{}, err error) {
	fmt.Println(id)
	fmt.Println(period)
	result := configs.DB.Unscoped().Debug().Where("member_id = ?", id).Delete(&models.Loan{})
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
