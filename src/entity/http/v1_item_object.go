package objects

import "time"

type V1ItemObjectResponse struct {
	ID               uint   `json:"id"`
	VendorID         int    `json:"vendor_id"`
	BrandID          int    `json:"brand_id"`
	ItemInitialID    int    `json:"item_initial_id"`
	Name             string `json:"name"`
	IsStockAvailable int    `json:"is_stock_available"`
	MinimumStock     int    `json:"minimum_stock"`
	MinimumOrder     int    `json:"minimum_order"`
}

type V1ItemObjectRequest struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
