package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sahilrush/src/models"
	"gorm.io/gorm"
)

type VideoController struct {
	DB *gorm.DB
}

// GetVideos retrieves paginated video records
func (vc *VideoController) GetVideos(c *gin.Context) {
	// Get pagination parameters from query string
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Calculate offset based on the current page
	offset := (page - 1) * limit

	// Query the database with pagination
	var videos []models.Video
	if err := vc.DB.Order("published_at desc").Offset(offset).Limit(limit).Find(&videos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with the videos in JSON format
	c.JSON(http.StatusOK, gin.H{
		"page":   page,
		"limit":  limit,
		"videos": videos,
	})
}
