package model

import (
	"errors"

	"gorm.io/gorm"
)

type Product struct {
	ID        int    `gorm:"autoIncrement;primary_key" json:"id"`
	Type      int    `json:"type"`
	Brand     string `json:"brand"`
	Series    string `json:"series"`
	Model     string `json:"model"`
	Color     string `json:"color"`
	Version   string `json:"version"`
	UpdatedAt string `json:"updated_at"`
	Prices    Prices `json:"prices" gorm:"foreignKey:ProductID"`
}

type Products []Product

type ProductType struct {
	ID   int    `gorm:"autoIncrement;primary_key" json:"id"`
	Name string `json:"name"`
}

type ProductTypes []ProductType

type Price struct {
	ID        int     `gorm:"autoIncrement;primary_key" json:"id"`
	ProductID int     `json:"product_id"`
	OutPrice  float64 `json:"out_price"`
	InPrice   float64 `json:"in_price"`
	Profit    float64 `json:"profit"`
	CreatedAt string  `json:"created_at"`
}

type Prices []Price

func (p *Products) ToIDs() []int {
	ids := make([]int, len(*p))
	for i, product := range *p {
		ids[i] = product.ID
	}
	return ids
}

func AddProduct(product Product) error {
	// 盘点是否存在相同的产品
	var p Product
	db.Where("type = ? AND brand = ? AND series = ? AND model = ? AND color = ? AND version = ?",
		product.Type, product.Brand, product.Series, product.Model, product.Color, product.Version).First(&p)
	if p.ID > 0 {
		return errors.New("product already exists")
	}
	db.Create(&product)
	return nil
}

func GetProductList(product_type int, brand string) (products Products) {
	db := db.Model(&Product{}).Preload("Prices", func(db *gorm.DB) *gorm.DB {
		return db.Order("id DESC")
	}).Where("type = ?", product_type)
	if brand != "" {
		db = db.Where("brand = ?", brand)
	}
	db.Find(&products)
	return
}

func GetPriceByDate(product_id int, date string) (price Price) {
	db.Where("product_id = ? AND created_at = ?", product_id, date).First(&price)
	return
}

func UpdatePrice(price Price) {
	db.Save(&price)
}

func AddPrice(price Price) {
	db.Create(&price)
}

func GetBrandList(t int) (brands []string) {
	db.Table("products").Select("brand").Where("type = ?", t).Group("brand").Find(&brands)
	return
}

func GetProductTypeList() (types ProductTypes) {
	db.Find(&types)
	return
}
