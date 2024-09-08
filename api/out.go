package api

type Out struct {
	ProductID   string `json:"product_id"`    // 产品ID
	OrderImg    string `json:"order_img"`     // 订单截图
	PkgFrontImg string `json:"pkg_front_img"` // 包装正面照片
	PkgBackImg  string `json:"pkg_back_img"`  // 包装背面照片
	Contact     string `json:"contact"`       // 联系方式
	Address     string `json:"address"`       // 地址
	SN          string `json:"sn"`            // 序列号
	Remark      string `json:"remark"`        // 备注
	Status      int    `json:"status"`        // 0 已创建 1 已完成
	CreateTime  string `json:"create_time"`   // 创建时间
	UpdateTime  string `json:"update_time"`   // 更新时间
}

type Outs []Out
