package controllers

import (
	"fmt"
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
	// Get pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}

	// Calculate offset
	offset := (page - 1) * limit
	fmt.Printf("Fetching videos with offset: %d, limit: %d\n", offset, limit)

	// Query the database
	var videos []models.Video
	if err := vc.DB.Order("published_at desc").Offset(offset).Limit(limit).Find(&videos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Log the videos for debugging
	fmt.Printf("Fetched videos: %+v\n", videos)

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"page":   page,
		"limit":  limit,
		"videos": videos,
	})
}
