package model

import "time"

type Product struct {
	ID        int     `gorm:"autoIncrement;primary_key" json:"id"`
	Brand     string  `json:"brand"`
	Series    string  `json:"series"`
	Model     string  `json:"model"`
	Color     string  `json:"color"`
	Version   string  `json:"version"`
	Price     float64 `json:"price" gorm:"-"`
	LastPrice float64 `json:"last_price" gorm:"-"`
}

type Products []Product

type Price struct {
	ID        int       `gorm:"autoIncrement;primary_key" json:"id"`
	ProductID int       `json:"product_id"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type Prices []Price

func (p *Products) ToIDs() []int {
	ids := make([]int, len(*p))
	for i, product := range *p {
		ids[i] = product.ID
	}
	return ids
}

func (p *Prices) GetTodayPrice(productId int) float64 {
	for _, p := range *p {
		if p.ProductID == productId && p.CreatedAt.Format("2006-01-02") == time.Now().Format("2006-01-02") {
			return p.Price
		}
	}
	return 0
}

func (p *Prices) GetLastDayPrice(productId int) float64 {
	for _, p := range *p {
		if p.ProductID == productId && p.CreatedAt.Format("2006-01-02") == time.Now().AddDate(0, 0, -1).Format("2006-01-02") {
			return p.Price
		}
	}
	return 0
}

func GetProductList(brand string) (products Products) {
	db := db.Model(&Product{})
	if brand != "" {
		db = db.Where("brand = ?", brand)
	}
	db.Find(&products)
	return
}

func GetPriceList(ids []int) (prices Prices) {
	db.Where("product_id IN (?)", ids).Find(&prices)
	return
}

func SavePrice(price Price) {
	var p Price
	// 根据 id 和 created_at(yyyy-mm-dd) 判断是否存在记录，如果存在，则更新，不存在则插入
	db.Where("product_id =? AND to_char(created_at, 'YYYY-MM-DD') = ?", price.ProductID, time.Now().Format("2006-01-02")).First(&p)
	if p.ID > 0 {
		db.Model(&p).Updates(Price{Price: price.Price})
	} else {
		db.Create(&price)
	}
}
