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
		return db.Order("created_at DESC")
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

// // 查询最后一次价格
// func GetLastPrice(productID int) (price Price) {
// 	db.Where("product_id = ? ", productID, time.Now()).Order("created_at desc").First(&price)
// 	return
// }

// func GetPriceList(ids []int) (prices Prices) {
// 	db.Where("product_id IN (?)", ids).Find(&prices)
// 	return
// }

// func GetCurrentPrice(productIDs []int) (prices Prices) {
// 	db.Where("product_id in (?) AND to_char(created_at, 'YYYY-MM-DD') = ?", productIDs, time.Now().Format("2006-01-02")).Find(&prices)
// 	return
// }

// func SavePrice(productID int, price float64, profit float64) {
// 	var p Price
// 	// 根据 id 和 created_at(yyyy-mm-dd) 判断是否存在记录，如果存在，则更新，不存在则插入
// 	db.Where("product_id =? AND to_char(created_at, 'YYYY-MM-DD') = ?", productID, time.Now().Format("2006-01-02")).First(&p)
// 	if p.ID > 0 {
// 		db.Model(&p).Updates(Price{RecoveryPrice: price + profit, Profit: profit, Price: price})
// 	} else {
// 		p.Price = price
// 		p.RecoveryPrice = price + profit
// 		p.ProductID = productID
// 		db.Create(&p)
// 	}
// }
