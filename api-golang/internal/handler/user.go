package handler

import (
	"log"
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

	if user.Username == "" || user.Password == "" || user.Email == "" || user.Role == "" {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Campos obrigatórios estão faltando"))
	}

	var existingUser models.User
	if err := config.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusConflict, utils.ErrorResponse("Username ou email já está em uso"))
	}

	validRoles := map[string]bool{"admin": true, "customer": true, "driver": true}
	if !validRoles[user.Role] {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Role inválido"))
	}

	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("SENHA INVALIDA"))
	}

	user.Password = password

	if err := config.DB.Create(&user).Error; err != nil {
		log.Println("Erro ao criar usuário:", err)
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

	var existingUser models.User
	if err := config.DB.First(&existingUser, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Usuário não encontrado"))
	}

	var updatedUser models.User
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Dados inválidos"))
	}

	if updatedUser.Password != "" {
		hashedPassword, err := utils.HashPassword(updatedUser.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Senha inválida"))
		}
		existingUser.Password = hashedPassword
	}

	existingUser.Username = updatedUser.Username
	existingUser.Email = updatedUser.Email
	existingUser.Role = updatedUser.Role
	existingUser.FullName = updatedUser.FullName
	existingUser.PhoneNumber = updatedUser.PhoneNumber

	if err := config.DB.Save(&existingUser).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Erro ao atualizar usuário"))
	}

	return c.JSON(http.StatusOK, existingUser)
}
