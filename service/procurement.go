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

func ProcurementCreate(dto *pojo.ProcurementCreateDTO) error {
	// 补全参数的默认值
	// 进货时间
	if dto.PurchaseTime == nil {
		dto.PurchaseTime = new(dao.Time)
		dto.PurchaseTime.Time = time.Now()
	}
	// 付款时间
	if dto.PayTime == nil {
		dto.PayTime = new(dao.Time)
		dto.PayTime.Time = time.Now()
	}
	// 实付金额
	totalPaid := decimal.Zero
	db := dao.GetDB()
	for i := range dto.Goods {
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
	procurement := dao.Procurement{
		TotalPaid:    *dto.TotalPaid,
		PayMethodID:  dto.PayMethodID,
		PurchaseTime: *dto.PurchaseTime,
		PayTime:      *dto.PayTime,
	}
	// 属性拷贝
	procurementGoods := make([]dao.ProcurementGoods, len(dto.Goods))
	for i := range dto.Goods {
		procurementGoods[i].GoodsID = dto.Goods[i].GoodsID
		procurementGoods[i].Quantity = dto.Goods[i].Quantity
		procurementGoods[i].Price = dto.Goods[i].Price
	}
	// 开启事务
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&procurement).Error; err != nil {
			return err
		}

		// 获取procurementId，并修改库存
		for i := range procurementGoods {
			// 添加订单id
			procurementGoods[i].ProcurementID = procurement.ID
			if *dto.ChangeStock {
				// 修改库存
				err := utils.ChangeStock(tx, procurementGoods[i].GoodsID, procurementGoods[i].Quantity)
				if err != nil {
					return err
				}
			}
		}

		if err := tx.Create(&procurementGoods).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func ProcurementPage(dto *pojo.ProcurementPageDTO) ([]dao.Procurement, int64, error) {
	db := dao.GetDB()

	tx := db.Model(&dao.Procurement{})
	// 查询条件
	if dto.PayMethodID != nil {
		tx = tx.Where("pay_method_id = ?", *dto.PayMethodID)
	}
	if dto.MinPurchaseTime != nil {
		tx = tx.Where("purchase_time >= ?", *dto.MinPurchaseTime)
	}
	if dto.MaxPurchaseTime != nil {
		tx = tx.Where("purchase_time <= ?", *dto.MaxPurchaseTime)
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
		tx = tx.Where("total_paid <= ?", *dto.MaxTotalPaid)
	}

	var total int64
	err := tx.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	var procurements []dao.Procurement
	offset := (dto.PageNum - 1) * dto.PageSize
	err = tx.Preload("PayMethod").Preload("ProcurementGoods").Order("id").Offset(offset).Limit(dto.PageSize).Find(&procurements).Error
	if err != nil {
		return nil, 0, err
	}
	return procurements, total, nil
}

func ProcurementFind(id int) (dao.Procurement, error) {
	db := dao.GetDB()
	var procurement dao.Procurement
	err := db.Preload("PayMethod").Preload("ProcurementGoods").Take(&procurement, id).Error
	return procurement, err
}

func ProcurementUpdate(procurement *dao.Procurement) error {
	db := dao.GetDB()
	err := db.Save(&procurement).Error
	return err
}

func ProcurementAddGoods(id int, dto []pojo.ProcurementAddUpdateGoodsDTO) error {
	db := dao.GetDB()

	err := db.Transaction(func(tx *gorm.DB) error {
		for i := range dto {
			// 先查找是否有记录
			procurementGoods := dao.ProcurementGoods{GoodsID: dto[i].GoodsID, ProcurementID: id}
			addQuantity := dto[i].Quantity
			goods := dao.Goods{ID: procurementGoods.GoodsID}
			if err := tx.Take(&procurementGoods).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				// 没有找到记录，添加记录
				if dto[i].Price == nil {
					// 从数据库查询单价
					if err := db.Select("price").Take(&goods).Error; err != nil {
						return err
					}
					procurementGoods.Price = goods.Price
				} else {
					// 直接赋值
					procurementGoods.Price = *dto[i].Price
				}
				procurementGoods.Quantity = dto[i].Quantity
			} else if err == nil {
				// 找到了记录，添加数量，但不会修改单价
				procurementGoods.Quantity = procurementGoods.Quantity.Add(addQuantity)
			} else {
				return err
			}
			err := tx.Save(&procurementGoods).Error
			if err != nil {
				return err
			}

			// 修改库存
			if dto[i].ChangeStock == nil || *dto[i].ChangeStock {
				err = utils.ChangeStock(tx, dto[i].GoodsID, addQuantity)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}

func ProcurementUpdateGoods(id int, dto []pojo.ProcurementAddUpdateGoodsDTO) error {
	db := dao.GetDB()

	err := db.Transaction(func(tx *gorm.DB) error {
		for i := range dto {
			procurementGoods := dao.ProcurementGoods{GoodsID: dto[i].GoodsID, ProcurementID: id}
			// 先查询原本的数量
			err := tx.Take(&procurementGoods).Error
			if err != nil {
				return err
			}
			addQuantity := dto[i].Quantity.Sub(procurementGoods.Quantity)

			// 判断是否删除订单中的商品
			if dto[i].Quantity.Equal(decimal.Zero) {
				err := tx.Delete(&procurementGoods).Error
				if err != nil {
					return err
				}

				// 修改库存
				if dto[i].ChangeStock == nil || *dto[i].ChangeStock {
					err = utils.ChangeStock(tx, dto[i].GoodsID, addQuantity)
					if err != nil {
						return err
					}
				}
				continue
			}

			goods := dao.Goods{ID: procurementGoods.GoodsID}
			// 修改数据
			procurementGoods.Quantity = dto[i].Quantity
			if dto[i].Price == nil {
				err := tx.Select("price").Take(&goods).Error
				if err != nil {
					return err
				}
				procurementGoods.Price = goods.Price
			} else {
				procurementGoods.Price = *dto[i].Price
			}
			err = tx.Save(&procurementGoods).Error
			if err != nil {
				return err
			}

			// 修改库存
			if dto[i].ChangeStock == nil || *dto[i].ChangeStock {
				err := utils.ChangeStock(tx, goods.ID, addQuantity)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}

func ProcurementDelete(id int, changeStock bool) error {
	db := dao.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		// 先修改库存
		if changeStock {
			// 查到商品数量
			var procurementGoods []dao.ProcurementGoods
			err := tx.Where("procurement_id = ?", id).Find(&procurementGoods).Error
			if err != nil {
				return err
			}

			for i := range procurementGoods {
				// 修改库存
				err := utils.ChangeStock(tx, procurementGoods[i].GoodsID, decimal.Zero.Sub(procurementGoods[i].Quantity))
				if err != nil {
					return err
				}
			}
		}

		// 删除连接表数据
		err := tx.Where("procurement_id = ?", id).Delete(&dao.ProcurementGoods{}).Error
		if err != nil {
			return err
		}

		// 删除订单
		err = tx.Delete(&dao.Procurement{ID: id}).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
