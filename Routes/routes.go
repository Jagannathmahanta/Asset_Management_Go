package Routes

import (
	"AssetManagementSystem/Controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	grp1 := r.Group("/ams/v1/")
	{
		grp1.POST("sign-in", Controllers.SignIn) //login logic
		grp1.POST("create-user", Controllers.CreateUser)
		grp1.GET("get-users", Controllers.FindUser)
		grp1.GET("get-assets", Controllers.FindAsset)
		grp1.POST("create-asset", Controllers.CreateAsset)
		grp1.PUT("update-user/:id", Controllers.Update)
		grp1.DELETE("delete-user/:id", Controllers.DeleteUsers)
		grp1.DELETE("deleteVendor/:id", Controllers.DeleteVendors)
		grp1.GET("get-tracking-by-asset/:asset_id", Controllers.FindTrackingByAsset)
		grp1.GET("get-total-assets-by-location/:location_id", Controllers.GetTotalAssetsByLocation)
		grp1.GET("get-owned-assets-by-location/:location_id", Controllers.GetOwnedAssetsByLocation)
		grp1.GET("get-rental-assets-by-location/:location_id", Controllers.GetRentalAssetsByLocation)
		grp1.GET("get-asset-details-by-locations", Controllers.GetAssetDetailsByLocations)

		grp1.GET("get-role", Controllers.FindRole)
		grp1.GET("get-location", Controllers.FindLocation)
		grp1.GET("get-owner", Controllers.FindOwnership)
		grp1.GET("get-asset-type", Controllers.FindAssetType)
		grp1.GET("get-model", Controllers.FindModels)
		grp1.GET("get-user/:id", Controllers.FindUserById)
		grp1.GET("get-vendor-By-location/:location", Controllers.FindVendorByLocation)
		grp1.DELETE("delete-asset/:id", Controllers.DeleteAsset)
		grp1.GET("get-manufacturer/:asset_type", Controllers.FindManufacturerByType)
		grp1.GET("get-asset-model/:manufacturer_name", Controllers.FindAssetModelByManufacturer)
		grp1.POST("assign-asset", Controllers.AddAssetForUsers)
		grp1.POST("create-model", Controllers.CreateModel)
		grp1.GET("get-available-asset", Controllers.GetAvailableAssets)
		grp1.DELETE("remove-assigned-asset", Controllers.DeleteAssetForUser)
		grp1.GET("get-vendor", Controllers.FindVendors)
		grp1.GET("logout", Controllers.Logout)
		grp1.POST("get-version", Controllers.FindVersionByModel)           //changed
		grp1.GET("assigned-asset/:user_id", Controllers.GetAssignedAssets) //added
		grp1.DELETE("delete-model/:id", Controllers.DeleteModel)
		grp1.POST("add-vendor", Controllers.CreateVendors)
		grp1.GET("get-assets-by-search/:input", Controllers.FindAssetDetailsByserialNo)
		//add vendor route
		grp1.POST("forgot-password", Controllers.ForgotPasswordSendUrl)
		grp1.POST("forgot-password/:user_id/:token", Controllers.ForgotPasswordNewPassword)
		grp1.POST("forget-password/:id", Controllers.ChangePassword)

	}
	return r
}
