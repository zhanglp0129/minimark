package utils

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"minimark/dao"
)

// ChangeStock 修改库存
func ChangeStock(tx *gorm.DB, goodsID int, addStock decimal.Decimal) error {
	goods := dao.Goods{ID: goodsID}
	// 先查到旧的库存
	err := tx.Select("stock").Take(&goods).Error
	if err != nil {
		return err
	}

	// 再增加库存
	err = tx.Model(&goods).Update("stock", goods.Stock.Add(addStock)).Error
	if err != nil {
		return err
	}
	return nil
}
