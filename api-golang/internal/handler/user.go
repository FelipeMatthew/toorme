package handler

import (
	"net/http"
	"strconv"
	"toorme-api-golang/config"
	"toorme-api-golang/internal/models"
	"toorme-api-golang/pkg/utils"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Dados inválidos"))
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Erro ao criar usuário"))
	}

	return c.JSON(http.StatusCreated, user)
}

func GetAllUser(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("error to fetch users"))
	}

	return c.JSON(http.StatusOK, users)
}

func GetUserById(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("ID inválido"))
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Usuário não encontrado"))
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("ID inválido"))
	}

	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Erro ao deletar usuário"))
	}

	return c.JSON(http.StatusOK, "USARIO APAGADO COM SUCESSO")
}

func UpdateUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("ID inválido"))
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Usuário não encontrado"))
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Dados inválidos"))
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Erro ao atualizar usuário"))
	}

	return c.JSON(http.StatusOK, user)
}
