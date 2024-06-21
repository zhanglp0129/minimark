package service

import (
	"errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"minimark/dao"
	"minimark/pojo"
	"minimark/utils"
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
				// 修改库存
				err := utils.ChangeStock(tx, orderGoods[i].GoodsID, decimal.Zero.Sub(orderGoods[i].Quantity))
				if err != nil {
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
		tx = tx.Where("pay_time >= ?", *dto.MinPayTime)
	}
	if dto.MaxPayTime != nil {
		tx = tx.Where("pay_time <= ?", *dto.MaxPayTime)
	}
	if dto.MinTotalPaid != nil {
		tx = tx.Where("total_paid >= ?", *dto.MinTotalPaid)
	}
	if dto.MaxTotalPaid != nil {
		tx = tx.Where("total_paid >= ?", *dto.MaxTotalPaid)
	}
	err := tx.Find(&orders).Error

	if err != nil {
		return nil, err
	}
	return orders, nil
}

func OrderFind(id int) (dao.Order, error) {
	db := dao.GetDB()
	var order dao.Order
	err := db.Preload("PayMethod").Preload("OrderGoods").Take(&order, id).Error
	return order, err
}

func OrderUpdate(order *dao.Order) error {
	db := dao.GetDB()
	err := db.Save(&order).Error
	return err
}

func OrderAddGoods(id int, dto []pojo.OrderAddUpdateGoodsDTO) error {
	db := dao.GetDB()

	err := db.Transaction(func(tx *gorm.DB) error {
		for i := range dto {
			// 先查找是否有记录
			orderGoods := dao.OrderGoods{GoodsID: dto[i].GoodsID, OrderID: id}
			addQuantity := dto[i].Quantity
			goods := dao.Goods{ID: orderGoods.GoodsID}
			if err := tx.Take(&orderGoods).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				// 没有找到记录，添加记录
				if dto[i].Price == nil {
					// 从数据库查询单价
					if err := db.Select("price").Take(&goods).Error; err != nil {
						return err
					}
					orderGoods.Price = goods.Price
				} else {
					// 直接赋值
					orderGoods.Price = *dto[i].Price
				}
				orderGoods.Quantity = dto[i].Quantity
			} else if err == nil {
				// 找到了记录，添加数量，但不会修改单价
				orderGoods.Quantity = orderGoods.Quantity.Add(addQuantity)
			} else {
				return err
			}
			err := tx.Save(&orderGoods).Error
			if err != nil {
				return err
			}

			// 修改库存
			if dto[i].ChangeStock == nil || *dto[i].ChangeStock {
				err = utils.ChangeStock(tx, dto[i].GoodsID, decimal.Zero.Sub(addQuantity))
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}

func OrderUpdateGoods(id int, dto []pojo.OrderAddUpdateGoodsDTO) error {
	db := dao.GetDB()

	err := db.Transaction(func(tx *gorm.DB) error {
		for i := range dto {
			orderGoods := dao.OrderGoods{GoodsID: dto[i].GoodsID, OrderID: id}
			// 先查询原本的数量
			err := tx.Take(&orderGoods).Error
			if err != nil {
				return err
			}
			addQuantity := dto[i].Quantity.Sub(orderGoods.Quantity)

			// 判断是否删除订单中的商品
			if dto[i].Quantity.Equal(decimal.Zero) {
				err := tx.Delete(&orderGoods).Error
				if err != nil {
					return err
				}

				// 修改库存
				if dto[i].ChangeStock == nil || *dto[i].ChangeStock {
					err = utils.ChangeStock(tx, dto[i].GoodsID, decimal.Zero.Sub(addQuantity))
					if err != nil {
						return err
					}
				}
				continue
			}

			goods := dao.Goods{ID: orderGoods.GoodsID}
			// 修改数据
			orderGoods.Quantity = dto[i].Quantity
			if dto[i].Price == nil {
				err := tx.Select("price").Take(&goods).Error
				if err != nil {
					return err
				}
				orderGoods.Price = goods.Price
			} else {
				orderGoods.Price = *dto[i].Price
			}
			err = tx.Save(&orderGoods).Error
			if err != nil {
				return err
			}

			// 修改库存
			if dto[i].ChangeStock == nil || *dto[i].ChangeStock {
				err := utils.ChangeStock(tx, goods.ID, decimal.Zero.Sub(addQuantity))
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}

func OrderDelete(id int, changeStock bool) error {
	db := dao.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		// 先修改库存
		if changeStock {
			// 查到商品数量
			var orderGoods []dao.OrderGoods
			err := tx.Where("order_id = ?", id).Find(&orderGoods).Error
			if err != nil {
				return err
			}

			for i := range orderGoods {
				// 修改库存
				err := utils.ChangeStock(tx, orderGoods[i].GoodsID, orderGoods[i].Quantity)
				if err != nil {
					return err
				}
			}
		}

		// 删除连接表数据
		err := tx.Where("order_id = ?", id).Delete(&dao.OrderGoods{}).Error
		if err != nil {
			return err
		}

		// 删除订单
		err = tx.Delete(&dao.Order{ID: id}).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
