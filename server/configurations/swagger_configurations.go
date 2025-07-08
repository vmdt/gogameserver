package configurations

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/vmdt/gogameserver/docs"
)

func ConfigSwagger(echo *echo.Echo) error {
	echo.GET("/swagger/*", echoSwagger.WrapHandler)
	return nil
}
