package pojo

import (
	"github.com/shopspring/decimal"
	"minimark/dao"
)

// GoodsPageDTO 接收条件分页查询前端传的参数
type GoodsPageDTO struct {
	PageNum    int     `form:"pageNum"`
	PageSize   int     `form:"pageSize"`
	CategoryID *int    `form:"categoryId"`
	GoodsName  *string `form:"goodsName"`
}

type GoodsInOrderProcurement struct {
	GoodsID  int              `form:"goodsId"`
	Quantity decimal.Decimal  `form:"quantity"`
	Price    *decimal.Decimal `form:"price"`
}

type OrderCreateDTO struct {
	TotalPaid   *decimal.Decimal          `form:"totalPaid"`
	PayMethodID *int                      `form:"payMethodId"`
	PayTime     *dao.Time                 `form:"payTime"`
	ChangeStock *bool                     `form:"changeStock"`
	Goods       []GoodsInOrderProcurement `form:"goods"`
}
