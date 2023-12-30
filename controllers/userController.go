package controllers

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"go-project/app"
	"go-project/database"
	"go-project/helpers"
	"go-project/models"
)

func Register(c *gin.Context) {
	var user app.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := database.DB.Where("email = ?", user.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusConflict, helpers.ErrorResponse{
			Success: false,
			Message: "Email already exist",
		})
		return
	}

	user.Password, _ = helpers.HashPassword(user.Password)

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, helpers.SuccesResponse{
		Success: true,
		Message: "Register Success",
		Data: app.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func Login(c *gin.Context) {

	var body app.LoginRequest
	var user app.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := database.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, helpers.ErrorResponse{
			Success: false,
			Message: "invalid password or email",
		})
		return
	}

	if err := helpers.CheckPasswordHash(body.Password, user.Password); err != true {
		c.JSON(http.StatusUnauthorized, helpers.ErrorResponse{
			Success: false,
			Message: "invalid password or email",
		})
		return
	}

	accessToken, err := helpers.GenerateToken(user.ID, user.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helpers.LoginResponse{
		Success:     true,
		Message:     "Login Success",
		AccessToken: accessToken,
		Data: app.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func UpdateUser(c *gin.Context) {

	userSession, exists := c.Get("user")

	userId, _ := c.Params.Get("userId")

	if !exists {
		c.JSON(http.StatusUnauthorized, helpers.ErrorResponse{
			Success: false,
			Message: "User not authenticated",
		})
		return
	}

	currentUser, ok := userSession.(*models.User)

	if !ok {
		c.JSON(http.StatusUnauthorized, helpers.ErrorResponse{
			Success: false,
			Message: "Failed to get user information",
		})
		return
	}

	if currentUser.ID != userId {
		c.JSON(http.StatusUnauthorized, helpers.ErrorResponse{
			Success: false,
			Message: "You are not allowed to update this user",
		})
		return
	}

	var user app.User

	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.ErrorResponse{
			Success: false,
			Message: "User not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user.Password, _ = helpers.HashPassword(user.Password)

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helpers.SuccesResponse{
		Success: true,
		Message: "Update Success",
		Data: app.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func DeleteUser(c *gin.Context) {

	currentUser, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, helpers.ErrorResponse{
			Success: false,
			Message: "User not authenticated",
		})
		return
	}

	user, ok := currentUser.(*models.User)

	if !ok {
		c.JSON(http.StatusUnauthorized, helpers.ErrorResponse{
			Success: false,
			Message: "Failed to get user information",
		})
		return
	}

	userId, _ := c.Params.Get("userId")

	if userId != user.ID {
		c.JSON(http.StatusForbidden, helpers.ErrorResponse{
			Success: false,
			Message: "You are not allowed to delete this user",
		})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helpers.SuccesResponse{
		Success: true,
		Message: "user deleted successfully",
	})
}
