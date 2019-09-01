package models

import (
	"time"
)

type Item struct {
	ID               uint       `gorm:"primary_key"`
	VendorID         uint       `gorm:"column:vendor_id" json:"vendor_id"`
	BrandID          uint       `gorm:"column:brand_id" json:"brand_id"`
	ItemInitialID    uint       `gorm:"column:item_initial_id" json:"item_initial_id"`
	Name             string     `gorm:"column:name" json:"name"`
	IsStockAvailable int        `gorm:"column:is_stock_available" json:"is_stock_available"`
	MinimumStock     int        `gorm:"column:minimum_stock" json:"minimum_stock"`
	MinimumOrder     int        `gorm:"column:minimum_order" json:"minimum_order"`
	CreatedAt        *time.Time `sql:"index"`
	UpdatedAt        *time.Time `sql:"index"`
	DeletedAt        *time.Time `sql:"index"`
}

func (Item) TableName() string {
	return "rl_items"
}
