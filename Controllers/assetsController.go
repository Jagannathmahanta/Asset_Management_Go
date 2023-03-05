package Controllers

import (
	"AssetManagementSystem/Config"
	"AssetManagementSystem/Models"
	"time"

	"github.com/gin-gonic/gin"

	//"database/sql"
	"net/http"
)

// API to create asset
func CreateAsset(c *gin.Context) {

	//validate input
	var input Models.CreateAsset
	var modelId Models.AssetModels
	//falsestatus := false
	truestatus := true
	currentDate := time.Now()

	// recdt := input.AssetReceivedDate.String()

	// recdate, err := time.Parse("2006-01-01", recdt)
	//fmt.Println(currentDate)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Config.DB.Table("asset_models").Select("id").Where("model_name = ? and version = ? and manufacturer_name= ?", input.ModelName, input.Version, input.ManufacturerName).First(&modelId)

	asset := Models.Assets{
		AssetModelName:    modelId.Id,
		AssetType:         input.AssetType,
		AssetOwner:        input.AssetOwner,
		Vendor:            input.Vendor,
		SerialNumber:      input.SerialNumber,
		AssetReceivedDate: &currentDate,
		AvailableStatus:   &truestatus,
		Configuration:     input.Configuration,
		Location:          input.Location,
		DeviceName:        input.DeviceName,
	}

	result := Config.DB.Create(&asset)

	//validate if unable to store data in db bez of foreign keys or duplicate entry
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Uable to create Asset... Please try again"})
		return
	}

	//success validation
	c.JSON(201,
		gin.H{
			"status":  "201",
			"message": "Assets created successfully",
		})
}

func FindAsset(c *gin.Context) {
	var asset []Models.Assetlist

	result := Config.DB.Model(&Models.Assets{}).
		Select("assets.id,asset_models.model_name, asset_types.asset_type, asset_ownerships.asset_owner, vendors.vendor_name, assets.serial_number, assets.asset_received_date, assets.asset_returned_date, assets.configuration,location_details.location, assets.device_name").
		Joins("JOIN asset_models on assets.asset_model_name = asset_models.id").
		Joins("JOIN asset_trackings on assets.id = asset_trackings.asset_id").
		Joins("JOIN users on asset_trackings.user_id = users.id").
		Joins("JOIN asset_types on assets.asset_type = asset_types.id").
		Joins("JOIN asset_ownerships on assets.asset_owner = asset_ownerships.id").
		Joins("LEFT JOIN vendors on assets.vendor = vendors.id").
		Joins("JOIN location_details on assets.location = location_details.id").
		Where("assets.asset_returned_date is null").
		Scan(&asset)

	//validation for not able get data
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to fetch Asset details"})
		return
	}
	//success validation

	c.JSON(http.StatusOK,
		gin.H{
			"status":  "200",
			"message": "Assets Fetched successfully",
			"Data":    asset,
		})
}

// GetTotalAssetsByLocation count value
func GetTotalAssetsByLocation(c *gin.Context) {

	var assets []Models.Assets
	var count int64

	result := Config.DB.Model(&assets).Where("location = ?", c.Param("location_id")).Count(&count)
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to fetch to total assets by location"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Total Assets Fetched successfully",
		"Count":   count})
}

// GetOwnedAssetsByLocation  count assets that are owned based on location
func GetOwnedAssetsByLocation(c *gin.Context) {

	var assets []Models.Assets
	var count int64

	result := Config.DB.Model(&assets).Where("location = ? and asset_owner=1", c.Param("location_id")).Count(&count)
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to fetch owned assets by location"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Owned assets Fetched successfully",
		"Count":   count})
}

// GetRentalAssetsByLocation count assets that are rented based on location
func GetRentalAssetsByLocation(c *gin.Context) {

	var assets []Models.Assets
	var count int64

	result := Config.DB.Model(&assets).Where("location = ? and asset_owner=2", c.Param("location_id")).Count(&count)
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to fetch rental assets by location"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Rental Assets Fetched successfully",
		"Count":   count})
}

func GetAssetDetailsByLocations(c *gin.Context) {
	var assets []Models.Assets
	var locations []Models.LocationDetails

	type Category struct {
		Location    string
		OwnedCount  int
		RentalCount int
		TotalCount  int
	}

	type Data struct {
		Categories []Category
	}

	var data Data

	var ownedCount int64
	var rentalCount int64
	var totalCount int64

	Config.DB.Find(&locations)

	for i := 0; i < len(locations); i++ {

		id := locations[i].Id
		Config.DB.Model(&assets).Where("location =? and asset_owner=1", id).Count(&ownedCount)
		Config.DB.Model(&assets).Where("location = ? and asset_owner=2", id).Count(&rentalCount)
		Config.DB.Model(&assets).Where("location = ?", id).Count(&totalCount)

		data.Categories = append(data.Categories, Category{
			Location:    locations[i].Location,
			TotalCount:  int(totalCount),
			OwnedCount:  int(ownedCount),
			RentalCount: int(rentalCount),
		})

	}
	if len(locations) == 0 {
		c.JSON(200,
			gin.H{
				"status":  "200",
				"message": "Asset doesn't exists",
			})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "data fetched successfully",
		"Data":    data.Categories,
	})

}

func GetAssignedAssets(c *gin.Context) {
	var assignAsset []Models.AssignedAssetlist

	//	result := Config.DB.Raw("select * from assets where id in (select asset_id from asset_trackings where user_id = ?)", c.Param("user_id")).Scan(&assignAsset)

	result := Config.DB.Model(&Models.Assets{}).
		Select("assets.id, asset_types.asset_type, asset_models.model_name, asset_models.version,assets.serial_number, vendors.vendor_name, location_details.location, asset_ownerships.asset_owner").
		Joins("JOIN asset_trackings on assets.id = asset_trackings.asset_id").
		Joins("JOIN asset_types on assets.asset_type = asset_types.id").
		Joins("JOIN asset_models on assets.asset_model_name = asset_models.id").
		Joins("LEFT JOIN vendors on assets.vendor = vendors.id").
		Joins("JOIN location_details on assets.location = location_details.id").
		Joins("JOIN asset_ownerships on assets.asset_owner = asset_ownerships.id").
		Where("user_id=? and asset_trackings.asset_returned_date is null", c.Param("user_id")).Scan(&assignAsset)

	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Uable to fetch assets"})
		return
	}
	if len(assignAsset) == 0 {
		c.JSON(http.StatusOK,
			gin.H{
				"status":  "200",
				"message": "No any assets assigned for this user",
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Assigned assets fetched successfully",
		"Data":    assignAsset})
}

// GetAvailableAssets get available assets
func GetAvailableAssets(c *gin.Context) {
	var assets []Models.Assetlist
	result := Config.DB.Model(&Models.Assets{}).
		Select("assets.id,asset_models.model_name, asset_types.asset_type, asset_ownerships.asset_owner, vendors.vendor_name, assets.serial_number, assets.asset_received_date, assets.asset_returned_date, assets.configuration,location_details.location").
		Joins("JOIN asset_models on assets.asset_model_name = asset_models.id").
		Joins("JOIN asset_types on assets.asset_type = asset_types.id").
		Joins("JOIN asset_ownerships on assets.asset_owner = asset_ownerships.id").
		Joins("LEFT JOIN vendors on assets.vendor = vendors.id").
		Joins("JOIN location_details on assets.location = location_details.id").
		Where("available_status=1").
		Scan(&assets)
	//validation for not able get data
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to fetch available assets"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Available Assets Fetched successfully",
		"Data":    assets})
}

// DeleteAsset API to Delete asset
func DeleteAsset(c *gin.Context) {
	var asset Models.Assets
	currentDate := time.Now()

	if err := Config.DB.Where("id=?", c.Param("id")).First(&asset).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	if err := Config.DB.Where("id=? and available_status=1", c.Param("id")).First(&asset).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Asset is assigned.. Please Remove assigned asset first!",
		})
		return
	}
	availableStatus := false
	//Config.DB.Find(&users)
	result := Config.DB.Model(&asset).Updates(Models.Assets{
		AvailableStatus: &availableStatus,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to delete asset... Please try again",
		})
		return
	}
	Config.DB.Model(&asset).Update(Models.Assets{
		AssetReturnedDate: &currentDate,
	})

	c.JSON(200, gin.H{
		"status":  204,
		"message": "Asset deleted successfully",
		"data":    asset,
	})
}
