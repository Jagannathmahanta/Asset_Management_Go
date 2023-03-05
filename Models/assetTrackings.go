package Models

import "time"

type AssetTracking struct {
	ID                uint       `json:"id"`
	AssetId           int        `json:"assetId"`
	UserId            int        `json:"userId"`
	AssetCondition    int        `json:"assetCondition"`
	AssetAssignedDate *time.Time `json:"assetAssignedDate"`
	AssetReturnedDate *time.Time `json:"assetReturnedDate"`
}
type AddAssertions struct {
	AssetId           int        `json:"assetId"`
	UserId            int        `json:"userId"`
	AssetCondition    int        `json:"assetCondition"`
	AssetAssignedDate *time.Time `json:"assetAssignedDate"`
	AssetReturnedDate *time.Time `json:"assetReturnedDate"`
}
