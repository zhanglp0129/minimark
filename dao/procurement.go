package dao

import "github.com/shopspring/decimal"

type Procurement struct {
	ID               int                `json:"id"`
	TotalPaid        decimal.Decimal    `gorm:"type:decimal(10,2);not null;check:total_paid>=0" json:"totalPaid"`
	PayMethodID      *int               `json:"payMethodId,omitempty"`
	PayMethod        *PayMethod         `json:"payMethod,omitempty"`
	PurchaseTime     Time               `gorm:"not null" json:"purchaseTime"`
	PayTime          Time               `gorm:"not null" json:"payTime"`
	Goods            []Goods            `gorm:"many2many:procurement_goods" json:"-"`
	ProcurementGoods []ProcurementGoods `json:"goods"`
}

type ProcurementGoods struct {
	GoodsID       int             `gorm:"primaryKey" json:"goodsId"`
	ProcurementID int             `gorm:"primaryKey" json:"procurementId"`
	Quantity      decimal.Decimal `gorm:"type:decimal(10,3);not null;check:quantity>=0" json:"quantity"`
	Price         decimal.Decimal `gorm:"type:decimal(10,2);not null;check:price>=0" json:"price"`
}
