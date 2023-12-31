package routes

import (
	"os"
	"prakerja12/controllers"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.Use(middleware.Logger())
	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)

	eAuth.POST("/member", controllers.CreateMemberController)
	eAuth.GET("/members", controllers.GetMembersController)

	eAuth.POST("/saving", controllers.CreateSavingController)
	eAuth.GET("/savings", controllers.GetSavingsController)
	eAuth.GET("/saving/:id", controllers.GetSavingByController)
	eAuth.DELETE("/saving", controllers.DeleteSavingByController)

	eAuth.POST("/loan", controllers.CreateLoanController)
	eAuth.GET("/loans", controllers.GetLoansController)
	eAuth.GET("/loan/:id", controllers.GetLoanByController)
	eAuth.DELETE("/loan", controllers.DeleteLoanByController)

	eAuth.GET("/loanpayment/:id", controllers.GetLoanPaymentByController)
}
