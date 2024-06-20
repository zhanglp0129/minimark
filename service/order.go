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
		totalPaid = totalPaid.Add(*dto.Goods[i].Price)
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
				var stock decimal.Decimal
				if err := db.Model(&dao.Goods{}).Select("stock").Take(&stock).Error; err != nil {
					return err
				}

				// 再减去购买数量
				if err := tx.Model(&dao.Goods{ID: orderGoods[i].GoodsID}).Update("stock", stock.Sub(orderGoods[i].Quantity)).Error; err != nil {
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
