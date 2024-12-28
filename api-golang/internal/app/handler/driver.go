package handler

import "github.com/labstack/echo/v4"

func DriverTrips(c echo.Context) error {
	return c.JSON(200, map[string]string{"message": "Aqui estão suas viagens designadas."})
}
