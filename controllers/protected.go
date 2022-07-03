package controllers

import (
	"azimio/database"
	"azimio/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// returns user data
func Profile(c *gin.Context) {
	var user models.User

	username, _ := c.Get("username") // from the authorization middleware

	result := database.DBCon.Where("username = ?", username.(string)).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"msg": "user not found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{
			"msg": "could not get user profile",
		})
		c.Abort()
		return
	}

	user.Password = ""

	c.JSON(200, user)

	return
}
