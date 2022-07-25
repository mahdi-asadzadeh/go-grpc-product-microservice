package models

import (
	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	Path string `gorm:"iniqueIndex"`
	Size int64 
	Product Product `gorm:"foreignKey:ProductID"`
	ProductID uint

}

type Product struct {
	gorm.Model
	Slug string `gorm:"uniqueIndex"`
	Title string `gorm:"size:2048"`
	Body string `gorm:"size:2048"`
	Price float64 
}
