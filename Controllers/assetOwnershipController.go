package Controllers

import (
	"AssetManagementSystem/Config"
	"AssetManagementSystem/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindOwnership API to get ownerships
func FindOwnership(c *gin.Context) {
	var ownership []Models.AssetOwnerships
	result := Config.DB.Find(&ownership)

	//validation for not able get data
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to fetch Owners"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Owners Fetched successfully",
		"Data":    ownership})
}
