package Models

type Vendors struct {
	ID           uint   `json:"id"`
	VendorName   string `'json:"vendorName"`
	MobileNumber int64  `json:"mobileNumber"`
	Email        string `json:"email"`
	Location     string `json:"location"`
	VendorStatus *bool  `json:"vendorStatus"`
}

type CreateVendor struct {
	VendorName   string `'json:"vendorName"`
	MobileNumber int64  `json:"mobileNumber"`
	Email        string `json:"email"`
	Location     int    `json:"location"`
	VendorStatus *bool  `json:"vendorStatus"`
}
