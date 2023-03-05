package Controllers

import (
	"AssetManagementSystem/Config"
	"AssetManagementSystem/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateVendors API to create Vendors
func CreateVendors(c *gin.Context) {
	var input Models.CreateVendor
	trueStatus := true
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vendor := Models.CreateVendor{

		VendorName:   input.VendorName,
		MobileNumber: input.MobileNumber,
		Email:        input.Email,
		Location:     input.Location,
		VendorStatus: &trueStatus,
	}
	result := Config.DB.Table("vendors").Create(&vendor)

	//validate if unable to store data in db bez of foreign keys or duplicate entry
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Unable to create Vendor... Please try again"})
		return
	}

	//success validation
	c.JSON(http.StatusCreated,
		gin.H{
			"status":  "201",
			"message": "Vendor created successfully",
		})
}

// FindVendors API to get Vendors
func FindVendors(c *gin.Context) {
	var vendors []Models.Vendors
	result := Config.DB.Table("vendors").
		Select("vendors.id, vendors.vendor_name, vendors.mobile_number, vendors.email, location_details.location,vendors.vendor_status").
		Joins("JOIN location_details on vendors.location = location_details.id").
		Where("vendor_status = 1").Scan(&vendors)

	//validation for not able get data
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to fetch"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Vendors Fetched successfully",
		"Data":    vendors})
}

// FindVendorByLocation API to get vendors by location
func FindVendorByLocation(c *gin.Context) {
	var vendorByLocation []Models.Vendors
	result := Config.DB.Table("vendors").Select("vendors.id, vendors.vendor_name, vendors.mobile_number, vendors.email, location_details.location,vendors.vendor_status").
		Joins("JOIN location_details on vendors.location = location_details.id").
		Where("vendor_status = 1 and location_details.id = ?", c.Param("location")).Scan(&vendorByLocation)

	//validation for not able get data
	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to fetch"})
		return
	}

	if len(vendorByLocation) == 0 {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Vendor's Not available for this location"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "vendors Fetched successfully",
		"Data":    vendorByLocation})
}

func DeleteVendors(c *gin.Context) {
	var vendor Models.Vendors

	if err := Config.DB.Where("id=? and vendor_status=1", c.Param("id")).First(&vendor).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}

	vendorStatus := false

	result := Config.DB.Model(&vendor).Updates(Models.Vendors{
		VendorStatus: &vendorStatus,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Unable to delete Vendor... Please try again"})
		return
	}

	c.JSON(200, gin.H{
		"status":  204,
		"message": "Vendor deleted successfully",
		"data":    vendor,
	})
}
