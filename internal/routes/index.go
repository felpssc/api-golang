package routes

import "github.com/labstack/echo/v4"

func Router(e *echo.Echo) {
	userRoutes(e)
}
