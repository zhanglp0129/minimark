package service

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"minimark/dao"
	"minimark/pojo"
	"time"
)

func OrderCreate(dto *pojo.OrderCreateDTO) error {
	// 补全参数的默认值
	// 付款时间
	if dto.PayTime == nil {
		dto.PayTime = new(dao.Time)
		dto.PayTime.Time = time.Now()
	}
	// 实付金额和购买时单价
	totalPaid := decimal.Zero
	db := dao.GetDB()
	for i := range dto.Goods {
		if dto.Goods[i].Price == nil {
			err := db.Model(&dao.Goods{}).Select("price").Find(&dto.Goods[i], dto.Goods[i].GoodsID).Error
			if err != nil {
				return err
			}
		}
		totalPaid = totalPaid.Add(dto.Goods[i].Price.Mul(dto.Goods[i].Quantity))
	}
	if dto.TotalPaid == nil {
		dto.TotalPaid = new(decimal.Decimal)
		*dto.TotalPaid = totalPaid
	}
	// 是否修改库存
	if dto.ChangeStock == nil {
		dto.ChangeStock = new(bool)
		*dto.ChangeStock = true
	}

	// 执行添加操作
	order := dao.Order{
		TotalPaid:   *dto.TotalPaid,
		PayMethodID: dto.PayMethodID,
		PayTime:     *dto.PayTime,
	}
	// 属性拷贝
	orderGoods := make([]dao.OrderGoods, len(dto.Goods))
	for i := range dto.Goods {
		orderGoods[i].GoodsID = dto.Goods[i].GoodsID
		orderGoods[i].Quantity = dto.Goods[i].Quantity
		orderGoods[i].Price = *dto.Goods[i].Price
	}
	// 开启事务
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		// 获取orderId，并修改库存
		for i := range orderGoods {
			// 添加订单id
			orderGoods[i].OrderID = order.ID
			if *dto.ChangeStock {
				// 先查到当前库存
				goods := dao.Goods{ID: orderGoods[i].GoodsID}
				if err := db.Select("stock").Take(&goods).Error; err != nil {
					return err
				}

				// 再减去购买数量
				if err := tx.Model(&dao.Goods{ID: orderGoods[i].GoodsID}).Update("stock", goods.Stock.Sub(orderGoods[i].Quantity)).Error; err != nil {
					return err
				}
			}
		}

		if err := tx.Create(&orderGoods).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func OrderPage(dto *pojo.OrderPageDTO) ([]dao.Order, error) {
	db := dao.GetDB()
	offset := (dto.PageNum - 1) * dto.PageSize

	// 查询Order
	var orders []dao.Order
	tx := db.Preload("PayMethod").Preload("OrderGoods").Offset(offset).Limit(dto.PageSize)
	// 查询条件
	if dto.PayMethodID != nil {
		tx = tx.Where("pay_method_id = ?", *dto.PayMethodID)
	}
	if dto.MinPayTime != nil {
		tx = tx.Where("pay_time >= ", *dto.MinPayTime)
	}
	if dto.MaxPayTime != nil {
		tx = tx.Where("pay_time <= ", *dto.MaxPayTime)
	}
	if dto.MinTotalPaid != nil {
		tx = tx.Where("total_paid >= ", *dto.MinTotalPaid)
	}
	if dto.MaxTotalPaid != nil {
		tx = tx.Where("total_paid >= ", *dto.MaxTotalPaid)
	}
	err := tx.Find(&orders).Error

	if err != nil {
		return nil, err
	}
	return orders, nil
}
