package pojo

import (
	"github.com/shopspring/decimal"
	"minimark/dao"
)

// GoodsPageDTO 接收条件分页查询前端传的参数
type GoodsPageDTO struct {
	PageNum    int     `form:"page_num"`
	PageSize   int     `form:"page_size"`
	CategoryID *int    `form:"category_id"`
	GoodsName  *string `form:"goods_name"`
}

type GoodsInOrder struct {
	GoodsID  int              `form:"goodsId"`
	Quantity decimal.Decimal  `form:"quantity"`
	Price    *decimal.Decimal `form:"price"`
}

type OrderCreateDTO struct {
	TotalPaid   *decimal.Decimal `form:"totalPaid"`
	PayMethodID *int             `form:"payMethodId"`
	PayTime     *dao.Time        `form:"payTime"`
	ChangeStock *bool            `form:"changeStock"`
	Goods       []GoodsInOrder   `form:"goods"`
}

type OrderPageDTO struct {
	PageNum      int              `form:"page_num"`
	PageSize     int              `form:"page_size"`
	PayMethodID  *int             `form:"pay_method_id"`
	MinPayTime   *dao.Time        `form:"min_pay_time"`
	MaxPayTime   *dao.Time        `form:"max_pay_time"`
	MinTotalPaid *decimal.Decimal `form:"min_total_paid"`
	MaxTotalPaid *decimal.Decimal `form:"max_total_paid"`
}

type OrderAddUpdateGoodsDTO struct {
	GoodsID     int              `json:"goodsId"`
	Quantity    decimal.Decimal  `json:"quantity"`
	Price       *decimal.Decimal `json:"price"`
	ChangeStock *bool            `json:"changeStock"`
}

type GoodsInProcurement struct {
	GoodsID  int             `form:"goodsId"`
	Quantity decimal.Decimal `form:"quantity"`
	Price    decimal.Decimal `form:"price"`
}

type ProcurementCreateDTO struct {
	TotalPaid    *decimal.Decimal     `form:"totalPaid"`
	PayMethodID  *int                 `form:"payMethodId"`
	PurchaseTime *dao.Time            `form:"purchaseTime"`
	PayTime      *dao.Time            `form:"payTime"`
	ChangeStock  *bool                `form:"changeStock"`
	Goods        []GoodsInProcurement `form:"goods"`
}

type ProcurementPageDTO struct {
	PageNum         int              `form:"page_num"`
	PageSize        int              `form:"page_size"`
	PayMethodID     *int             `form:"pay_method_id"`
	MinPurchaseTime *dao.Time        `form:"min_purchase_time"`
	MaxPurchaseTime *dao.Time        `form:"max_purchase_time"`
	MinPayTime      *dao.Time        `form:"min_pay_time"`
	MaxPayTime      *dao.Time        `form:"max_pay_time"`
	MinTotalPaid    *decimal.Decimal `form:"min_total_paid"`
	MaxTotalPaid    *decimal.Decimal `form:"max_total_paid"`
}

type ProcurementAddUpdateGoodsDTO struct {
	GoodsID     int              `json:"goodsId"`
	Quantity    decimal.Decimal  `json:"quantity"`
	Price       *decimal.Decimal `json:"price"`
	ChangeStock *bool            `json:"changeStock"`
}
