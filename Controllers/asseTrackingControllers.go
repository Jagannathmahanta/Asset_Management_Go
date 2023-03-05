package Controllers

import (
	"AssetManagementSystem/Config"
	"AssetManagementSystem/Models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// FindTrackingByAsset get tracking of a particular asset by id
func FindTrackingByAsset(c *gin.Context) {
	var tracking []Models.AssetTracking

	result := Config.DB.Where("asset_id = ?", c.Param("asset_id")).Find(&tracking) //use First if only single record need to be found
	//validation for not able get data
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to fetch asset details"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "asset details Fetched successfully",
		"Data":    tracking})
}

// AddAssetForUsers API to add asset for user
func AddAssetForUsers(c *gin.Context) {
	currentDate := time.Now()
	//variable with boolean value
	falseStatus := false

	var updateStatus []Models.Assets

	//validate input
	var input Models.AssetTracking

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Checks if Asset is available or not
	Config.DB.Table("assets").Where("id=? and available_status=1", input.AssetId).First(&updateStatus)
	if len(updateStatus) == 0 {
		c.JSON(http.StatusOK,
			gin.H{
				"status":  "200",
				"message": "Asset is not available",
			})
		return
	}

	//if asset is not assigned, This will add asset for user
	addAsset := Models.AddAssertions{
		AssetId:           input.AssetId,
		UserId:            input.UserId,
		AssetCondition:    1,
		AssetAssignedDate: &currentDate,
	}
	result := Config.DB.Table("asset_trackings").Create(&addAsset)
	//validate if unable to store data
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to add Asset for user... Please try again"})
		return
	}

	//this is to change asset status to unavailable
	Config.DB.Model(&updateStatus).Where("id=?", input.AssetId).Updates(Models.Assets{
		AvailableStatus: &falseStatus,
	})

	//When Data Stored successfully
	c.JSON(200,
		gin.H{
			"status":  "201",
			"message": "Asset added for user successfully",
			//"Data":    addAsset,
		})
}

// DeleteAssetForUser delete assets for users
func DeleteAssetForUser(c *gin.Context) {

	//variable with true value
	trueStatus := true
	var assetForUser []Models.AssetTracking
	var updateStatus []Models.Assets
	currentDate := time.Now()
	//validate input
	var input Models.AssetTracking

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	Config.DB.Table("asset_trackings").Where("asset_id = ? and user_id = ? and asset_returned_date is null", input.AssetId, input.UserId).Find(&assetForUser)

	if len(assetForUser) != 0 {
		if (Config.DB.Model(&assetForUser).Where("asset_id = ? and user_id = ?", input.AssetId, input.UserId).
			Updates(Models.AssetTracking{AssetReturnedDate: &currentDate})).Error != nil {
			c.JSON(200,
				gin.H{
					"status":  "200",
					"message": "Unable to delete asset for user",
				})
			return

		}
		//this is to change asset status to available
		Config.DB.Model(&updateStatus).Where("id=?", input.AssetId).Updates(Models.Assets{
			AvailableStatus: &trueStatus,
		})
		c.JSON(200,
			gin.H{
				"status":  "204",
				"message": "Asset has deleted for user",
			})

	} else {
		c.JSON(404,
			gin.H{
				"status":  "404",
				"message": "Record not found",
			})
	}

}

func FindAssetDetailsByserialNo(c *gin.Context) {
	var asset []Models.Assetlist2
	var asset1 []Models.Assetlist1

	Config.DB.Model(&Models.Assets{}).Where("serial_number=?", c.Param("input")).Scan(&asset1)
	Config.DB.Model(&Models.Assets{}).Where("device_name=?", c.Param("input")).Scan(&asset)

	if len(asset1) != 0 {
		result := Config.DB.Model(&Models.Assets{}).
			Select("assets.id,asset_models.model_name,asset_models.version,asset_models.manufacturer_name, asset_types.asset_type, asset_ownerships.asset_owner, vendors.vendor_name,available_status, assets.configuration,location_details.location, assets.device_name,users.username,asset_trackings.asset_assigned_date").
			Joins("JOIN asset_models on assets.asset_model_name = asset_models.id").
			Joins("LEFT JOIN asset_trackings on assets.id = asset_trackings.asset_id").
			Joins("LEFT JOIN users on asset_trackings.user_id = users.id").
			Joins("JOIN asset_types on assets.asset_type = asset_types.id").
			Joins("JOIN asset_ownerships on assets.asset_owner = asset_ownerships.id").
			Joins("LEFT JOIN vendors on assets.vendor = vendors.id").
			Joins("JOIN location_details on assets.location = location_details.id").
			Where("assets.serial_number=?  and assets.asset_returned_date is NULL ", c.Param("input")).
			Order("asset_trackings.asset_returned_date").
			Limit(1).
			Scan(&asset1)

		//validation for not able get data
		if result.Error != nil {
			c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to fetch Asset details"})
			return
		}

		c.JSON(200,
			gin.H{
				"status":  "200",
				"message": "Assets Fetched successfully",
				"Data":    asset1,
			})
	} else if len(asset) != 0 {
		result := Config.DB.Model(&Models.Assets{}).
			Select("assets.id,asset_models.model_name,asset_models.version,asset_models.manufacturer_name, asset_types.asset_type, asset_ownerships.asset_owner, vendors.vendor_name, assets.serial_number ,available_status,  assets.configuration,location_details.location,users.username,asset_trackings.asset_assigned_date").
			Joins("JOIN asset_models on assets.asset_model_name = asset_models.id").
			Joins("LEFT JOIN asset_trackings on assets.id = asset_trackings.asset_id").
			Joins("LEFT JOIN users on asset_trackings.user_id = users.id").
			Joins("JOIN asset_types on assets.asset_type = asset_types.id").
			Joins("JOIN asset_ownerships on assets.asset_owner = asset_ownerships.id").
			Joins("LEFT JOIN vendors on assets.vendor = vendors.id").
			Joins("JOIN location_details on assets.location = location_details.id").
			Where("assets.device_name=?  and assets.asset_returned_date is NULL ", c.Param("input")).
			Order("asset_trackings.asset_returned_date").
			Limit(1).
			Scan(&asset)

		//validation for not able get data
		if result.Error != nil {
			c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to fetch Asset details"})
			return
		}

		c.JSON(200,
			gin.H{
				"status":  "200",
				"message": "Assets Fetched successfully",
				"Data":    asset,
			})
	} else {

		c.JSON(200,
			gin.H{
				"status":  "400",
				"message": "Asset doesn't exists",
			})
	}

}
