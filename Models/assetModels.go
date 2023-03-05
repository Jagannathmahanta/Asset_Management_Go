package Models

type AssetModels struct {
	Id               int    `json:"id"`
	ModelName        string `json:"modelName"`
	Version          string `json:"version"`
	AssetType        int    `json:"assetType"`
	ModelStatus      *bool  `json:"modelStatus"`
	ManufacturerName string `json:"manufacturerName"`
}

type DeleteModel struct {
	ModelName string `json:"modelName"`
	Version   string `json:"version"`
	AssetType string `json:"assetType"`
}

type ModelByVer struct {
	Id      int    `json:"id"`
	Version string `json:"version"`
}

type ManufacturerByType struct {
	ManufacturerName string `json:"manufacturerName"`
}
type ModelByType struct {
	Id        int    `json:"id"`
	ModelName string `json:"modelName"`
	AssetType int    `json:"assetType"`
}

type ModelName struct {
	ManufacturerName string `json:"ManufacturerName"`
	ModelName        string `json:"modelName"`
	AssetType        string `json:"assetType"`
	Id               string `json:"id"`
}
