package Controllers

import (
	"AssetManagementSystem/Config"
	"AssetManagementSystem/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindAssetType API to get asset types
func FindAssetType(c *gin.Context) {
	var assetType []Models.AssetTypes
	result := Config.DB.Find(&assetType)

	//validation for not able get data
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to fetch asset types",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Asset types fetched successfully",
		"Data":    assetType})
}
