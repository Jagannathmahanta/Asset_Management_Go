package Controllers

import (
	"AssetManagementSystem/Config"
	"AssetManagementSystem/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteModel API to delete model
func DeleteModel(c *gin.Context) {
	// API to Delete asset
	var model Models.AssetModels

	falseStatus := false

	if err := Config.DB.Where("id = ? and model_status=1", c.Param("id")).
		First(&model).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	result := Config.DB.Model(&model).Where("id = ?", c.Param("id")).
		Updates(Models.AssetModels{
			ModelStatus: &falseStatus,
		})

	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to delete Model... Please try again",
		})
		return
	}

	c.JSON(201, gin.H{
		"status":  201,
		"message": "Model deleted successfully",
		"data":    model,
	})
}

// CreateModel API to create Model
func CreateModel(c *gin.Context) {
	modelStatus := true
	//validate input
	var input Models.AssetModels
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addModel := Models.AssetModels{
		ManufacturerName: input.ManufacturerName,
		ModelName:        input.ModelName,
		Version:          input.Version,
		AssetType:        input.AssetType,
		ModelStatus:      &modelStatus,
	}

	result := Config.DB.Create(&addModel)

	//validate if unable to store data in db bez of foreign keys or duplicate entry
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to create Model... Please try again"})
		return
	}

	//success validation
	c.JSON(http.StatusCreated,
		gin.H{
			"status":  "201",
			"message": "Model created successfully",
		})
}

// FindModels API to get models
func FindModels(c *gin.Context) {
	var models []Models.ModelName
	result := Config.DB.Raw("select Distinct asset_models.manufacturer_name, asset_models.model_name, asset_models.id,asset_types.asset_type from asset_models INNER JOIN asset_types ON asset_models.asset_type = asset_types.id where model_status = 1").Scan(&models)

	//validation for not able get data
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to fetch Models"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Models Fetched successfully",
		"Data":    models})
}

func FindManufacturerByType(c *gin.Context) {
	var manufacturer []Models.ManufacturerByType

	result := Config.DB.Raw("SELECT distinct manufacturer_name FROM asset_models INNER JOIN asset_types ON asset_models.asset_type = asset_types.id where asset_types.asset_type=?", c.Param("asset_type")).Scan(&manufacturer)

	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to fetch Manufacturer's by asset type"})
		return
	}
	if len(manufacturer) == 0 {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Manufacturer NOT available for this asset type"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Manufacturer's Fetched successfully",
		"Data":    manufacturer})
}

// FindAssetModelByAssetType API to get asset models by asset type
func FindAssetModelByManufacturer(c *gin.Context) {
	var assetModelByManufacturer []Models.ModelByType
	result := Config.DB.Table("asset_models").
		Where("manufacturer_name = ? and model_status=1", c.Param("manufacturer_name")).
		Scan(&assetModelByManufacturer)
	//validation for not able get data
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to fetch asset models by Manufacturer"})
		return
	}
	if len(assetModelByManufacturer) == 0 {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Asset model's Not available for this Manufacturer"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Asset Models Fetched successfully",
		"Data":    assetModelByManufacturer})
}

// FindVersionByModel API to get version By models
func FindVersionByModel(c *gin.Context) {
	var versionByModel []Models.ModelByVer
	var input Models.ModelName
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := Config.DB.Table("asset_models").
		Where("model_name = ? and manufacturer_name=? and model_status=1", input.ModelName, input.ManufacturerName).
		Scan(&versionByModel)
	//validation for not able get data
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to fetch version by models"})
		return
	}
	if len(versionByModel) == 0 {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Version's Not available for this model"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Versions Fetched successfully",
		"Data":    versionByModel})
}
