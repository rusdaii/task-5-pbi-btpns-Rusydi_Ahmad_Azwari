package controllers

import (
	"go-project/app"
	"go-project/database"
	"go-project/helpers"
	"go-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {

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

	var photo app.Photo

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	photo.UserID = user.ID

	if err := database.DB.Create(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, helpers.SuccesResponse{
		Success: true,
		Message: "Create Photo Successfully",
		Data:    photo,
	})
}

func GetAllPhotos(c *gin.Context) {

	var photos []app.Photo

	if err := database.DB.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helpers.SuccesResponse{
		Success: true,
		Message: "Get Photos Success",
		Data:    photos,
	})
}

func GetPhotoById(c *gin.Context) {

	photoId, _ := c.Params.Get("photoId")

	var photo app.Photo

	if err := database.DB.Where("id = ?", photoId).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.ErrorResponse{
			Success: false,
			Message: "Photo not found",
		})
		return
	}

	c.JSON(http.StatusOK, helpers.SuccesResponse{
		Success: true,
		Message: "Get Photo Success",
		Data:    photo,
	})
}

func UpdatePhoto(c *gin.Context) {

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

	photoId, _ := c.Params.Get("photoId")

	var photo app.Photo

	if err := database.DB.Where("id = ?", photoId).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.ErrorResponse{
			Success: false,
			Message: "Photo not found",
		})
		return
	}

	if photo.UserID != user.ID {
		c.JSON(http.StatusForbidden, helpers.ErrorResponse{
			Success: false,
			Message: "You are not allowed to update this photo",
		})
		return
	}

	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if err := database.DB.Save(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helpers.SuccesResponse{
		Success: true,
		Message: "Update Photo Successfully",
		Data:    photo,
	})
}

func DeletePhoto(c *gin.Context) {

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

	photoId, _ := c.Params.Get("photoId")

	var photo app.Photo

	if err := database.DB.Where("id = ?", photoId).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, helpers.ErrorResponse{
			Success: false,
			Message: "Photo not found",
		})
		return
	}

	if photo.UserID != user.ID {
		c.JSON(http.StatusForbidden, helpers.ErrorResponse{
			Success: false,
			Message: "You are not allowed to delete this photo",
		})
		return
	}

	if err := database.DB.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helpers.ErrorResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helpers.SuccesResponse{
		Success: true,
		Message: "Delete Photo Successfully",
	})
}
