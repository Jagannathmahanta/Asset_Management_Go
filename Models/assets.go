package Models

import "time"

type MyTimee time.Time

type Assets struct {
	ID                uint       `json:"ID"`
	AssetModelName    int        `json:"assetModelName"`
	AssetType         int        `json:"assetType"`
	AssetOwner        int        `json:"assetOwner"`
	Vendor            int        `json:"vendor"`
	SerialNumber      string     `json:"serialNumber"`
	AssetReceivedDate *time.Time `json:"assetReceivedDate"`
	AssetReturnedDate *time.Time `json:"assetReturnedDate"`
	AvailableStatus   *bool      `json:"availableStatus"`
	Configuration     string     `json:"configuration"`
	Location          int        `json:"location"`
	DeviceName        string     `json:"deviceName"`
}

type Assetlist struct {
	Id            int        `json:"id"`
	ModelName     string     `json:"modelName"`
	AssetType     string     `json:"assetType"`
	AssetOwner    string     `json:"assetOwner"`
	VendorName    string     `json:"vendorName"`
	SerialNumber  string     `json:"serialNumber"`
	ReceivedDate  *time.Time `json:"receivedDate"`
	ReturnedDate  *time.Time `json:"returnedDate"`
	Configuration string     `json:"configuration"`
	Location      string     `json:"location"`
	DeviceName    string     `json:"deviceName"`
}

type CreateAsset struct {
	ModelName         string     `json:"modelName"`
	Version           string     `json:"version"`
	AssetType         int        `json:"assetType"`
	ManufacturerName  string     `json:"manufacturerName"`
	AssetOwner        int        `json:"assetOwner"`
	Vendor            int        `json:"vendor"`
	SerialNumber      string     `json:"serialNumber"`
	AssetReceivedDate *time.Time `json:"assetReceivedDate"`
	AssetReturnedDate *time.Time `json:"assetReturnedDate"`
	AvailableStatus   *bool      `json:"availableStatus"`
	Configuration     string     `json:"configuration"`
	Location          int        `json:"location"`
	DeviceName        string     `json:"deviceName"`
}

type Data struct {
	Count AssetsCount
}
type AssetsCount struct {
	//location    string
	TotalCount  int
	OwnedCount  int
	RentalCount int
}

type AssignedAssetlist struct {
	Id           int    `json:"id"`
	ModelName    string `json:"modelName"`
	AssetType    string `json:"assetType"`
	VendorName   string `json:"vendorName"`
	AssetOwner   string `json:"assetOwner"`
	Version      string `json:"version"`
	Location     string `json:"location"`
	SerialNumber string `json:"serialNumber"`
}

type Assetlist2 struct {
	Id                int       `json:"id"`
	ModelName         string    `json:"modelName"`
	Version           string    `json:"version"`
	ManufacturerName  string    `json:"manufacturerName"`
	AssetType         string    `json:"assetType"`
	AssetOwner        string    `json:"assetOwner"`
	VendorName        string    `json:"vendorName"`
	SerialNumber      string    `json:"serialNumber"`
	AvailableStatus   *bool     `json:"availableStatus"`
	Configuration     string    `json:"configuration"`
	Location          string    `json:"location"`
	Username          string    `json:"username"`
	AssetAssignedDate time.Time `json:"assetAssignedDate"`
}
type Assetlist1 struct {
	Id                int       `json:"id"`
	ModelName         string    `json:"modelName"`
	Version           string    `json:"version"`
	ManufacturerName  string    `json:"manufacturerName"`
	AssetType         string    `json:"assetType"`
	AssetOwner        string    `json:"assetOwner"`
	VendorName        string    `json:"vendorName"`
	AvailableStatus   *bool     `json:"availableStatus"`
	Configuration     string    `json:"configuration"`
	Location          string    `json:"location"`
	DeviceName        string    `json:"deviceName"`
	Username          string    `json:"username"`
	AssetAssignedDate time.Time `json:"assetAssignedDate"`
}
