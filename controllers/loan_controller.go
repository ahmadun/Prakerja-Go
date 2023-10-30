package controllers

import (
	"fmt"
	"net/http"
	"prakerja12/models"
	"prakerja12/repositories"

	"github.com/labstack/echo/v4"
)

func CreateLoanController(c echo.Context) error {
	var req models.Loan
	c.Bind(&req)

	err := repositories.AddLoan(&req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: "Berhasil menambah data",
		Status:  true,
		Data:    req,
	})
}

func GetLoansController(c echo.Context) error {
	var res []models.Loan

	err := repositories.GetLoans(&res)
	fmt.Println(res)

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

func GetLoanByController(c echo.Context) error {
	var res []models.Loan
	id := c.Param("id")
	err := repositories.GetLoanBy(&res, id)

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

func DeleteLoanByController(c echo.Context) error {
	id := c.QueryParam("id")
	period := c.QueryParam("period")
	res, err := repositories.DelLoan(id, period)

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
