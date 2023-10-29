package controllers

import (
	"fmt"
	"net/http"
	"prakerja12/models"
	"prakerja12/repositories"

	"github.com/labstack/echo/v4"
)

func CreateMemberController(c echo.Context) error {
	var req models.Member
	c.Bind(&req)

	err := repositories.AddMember(&req)

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

func GetMembersController(c echo.Context) error {
	var res []models.Member

	err := repositories.GetMembers(&res)
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
