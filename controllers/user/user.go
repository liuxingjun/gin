package user

import (
	"gin/lib"
	"gin/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, lib.Gorm.First(&models.User{}))
}
