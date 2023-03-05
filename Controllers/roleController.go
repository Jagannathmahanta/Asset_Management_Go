package Controllers

import (
	"AssetManagementSystem/Config"
	"AssetManagementSystem/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindRole API to get roles
func FindRole(c *gin.Context) {
	var roles []Models.Roles
	result := Config.DB.Find(&roles)

	//validation for not able get data
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to fetch roles",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Roles Fetched successfully",
		"Data":    roles})
}
