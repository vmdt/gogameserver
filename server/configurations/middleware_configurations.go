package configurations

import (
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(recoverWithStackTrace(os.Getenv("APP_ENV")))
}

func recoverWithStackTrace(environment string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					stack := debug.Stack()
					errMsg := fmt.Sprintf("Recovered from panic: %v\nStack trace:\n%s", err, stack)
					c.Logger().Error(errMsg)

					if !c.Response().Committed {
						if environment == "development" {
							c.JSON(http.StatusInternalServerError, map[string]interface{}{
								"error": fmt.Sprintf("%v", err),
								"stack": string(stack),
							})
						} else {
							c.JSON(http.StatusInternalServerError, map[string]string{
								"error":   fmt.Sprintf("%v", err),
								"message": "An unexpected error occurred. Please try again later.",
							})
						}

					}
				}
			}()
			return next(c)
		}
	}
}
