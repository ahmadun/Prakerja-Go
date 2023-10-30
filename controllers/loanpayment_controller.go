package controllers

import (
	"net/http"
	"prakerja12/models"
	"prakerja12/repositories"

	"github.com/labstack/echo/v4"
)

func GetLoanPaymentByController(c echo.Context) error {
	var res []models.LoanPayment
	id := c.Param("id")
	err := repositories.GetLoanPaymentBy(&res, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menampilkan data",
		Status:  true,
		Data:    res,
	})
}
