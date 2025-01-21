package handler

import (
	"net/http"
	"toorme-api-golang/config"
	"toorme-api-golang/internal/models"

	"github.com/labstack/echo/v4"
)

func FetchAllData(c echo.Context) error {
	var users []models.User
	var suppliers []models.Supplier
	var vehicles []models.Vehicle
	var trips []models.Trip
	var travelPlans []models.TravelPlan
	var locations []models.Location

	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
	}

	if err := config.DB.Find(&suppliers).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch suppliers"})
	}

	if err := config.DB.Find(&vehicles).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch vehicles"})
	}

	if err := config.DB.Find(&trips).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch trips"})
	}

	if err := config.DB.Find(&travelPlans).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch travel plans"})
	}

	if err := config.DB.Find(&locations).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch locations"})
	}

	response := map[string]interface{}{
		"users":        users,
		"suppliers":    suppliers,
		"vehicles":     vehicles,
		"trips":        trips,
		"travel_plans": travelPlans,
		"locations":    locations,
	}

	return c.JSON(http.StatusOK, response)

}
