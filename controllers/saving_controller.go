package controllers

import (
	"fmt"
	"net/http"
	"prakerja12/models"
	"prakerja12/repositories"

	"github.com/labstack/echo/v4"
)

func CreateSavingController(c echo.Context) error {
	var req models.Saving
	c.Bind(&req)

	err := repositories.AddSaving(&req)

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

func GetSavingsController(c echo.Context) error {
	var res []models.Saving

	err := repositories.GetSavings(&res)
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
