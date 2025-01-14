package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllCustumer(c echo.Context) error {

	return c.JSON(http.StatusOK, map[string]string{"message": "Aqui estão seus planos de viagem."})
}
