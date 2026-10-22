package middlewares

import (
	"fmt"

	"github.com/StackStoreHQ/go-fm/modules"
	"github.com/labstack/echo/v4"
)

func LoggerIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		modules.Logger.Info(fmt.Sprintf(">>> %s %s", c.Request().Method, c.Request().URL.String()))
		return next(c)
	}
}
