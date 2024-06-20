package pojo

import "minimark/dao"

type OrderPageVO struct {
	dao.Order
	Goods []dao.OrderGoods `json:"goods"`
}
