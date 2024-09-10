package model

import "time"

type Sale struct {
	ID          int       `gorm:"autoIncrement;primary_key" json:"id"`
	ProductID   int       `json:"product_id"`    // 产品ID
	UserID      int       `json:"user_id"`       // 用户ID
	OrderImg    string    `json:"order_img"`     // 订单截图
	PkgFrontImg string    `json:"pkg_front_img"` // 包装正面照片
	PkgBackImg  string    `json:"pkg_back_img"`  // 包装背面照片
	Contact     string    `json:"contact"`       // 联系方式
	Address     string    `json:"address"`       // 地址
	SN          string    `json:"sn"`            // 序列号
	Remark      string    `json:"remark"`        // 备注
	Status      int       `json:"status"`        // 0 已创建 1 已完成
	CreateTime  time.Time `json:"create_time"`   // 创建时间
	UpdateTime  time.Time `json:"update_time"`   // 更新时间
	Price       float64   `json:"price"`         // 价格
	Receiver    int       `json:"receiver"`      // 收货人
	ReceiveTime string    `json:"receive_time"`  // 收货时间
	Payment     int       `json:"payment"`       // 支付方式
	Product     Product   `json:"product" gorm:"foreignKey:ProductID"`
}

type Sales []Sale

func CreateSale(s Sale) (err error) {
	s.CreateTime = time.Now()
	return db.Create(&s).Error
}

func GetSale(id int) (sale Sale, err error) {
	err = db.Preload("Product").First(&sale, id).Error
	return
}

func GetSales(userID, status int) (sales Sales, err error) {
	db := db.Model(&Sale{}).Where("status = ?", status)
	if userID > 0 {
		db = db.Where("user_id = ?", userID)
	}
	err = db.Preload("Product").Find(&sales).Error
	return
}

func GetSalesByUser(userID int) (sales Sales, err error) {
	err = db.Where("user_id = ?", userID).Find(&sales).Error
	return
}

func ConfirmSale(id int, price float64, payment int, receiver int) (err error) {
	sale := Sale{}
	err = db.First(&sale, id).Error
	if err != nil {
		return
	}
	sale.Price = price
	sale.Status = 1
	sale.Payment = payment
	sale.Receiver = receiver
	sale.UpdateTime = time.Now()
	err = db.Save(&sale).Error
	return
}
