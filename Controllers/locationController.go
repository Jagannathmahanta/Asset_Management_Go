package Controllers

import (
	"AssetManagementSystem/Config"
	"AssetManagementSystem/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindLocation API to get locations
func FindLocation(c *gin.Context) {
	var locations []Models.LocationDetails
	result := Config.DB.Find(&locations)

	//validation for not able get data
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to fetch locations",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Locations Fetched successfully",
		"Data":    locations})
}
