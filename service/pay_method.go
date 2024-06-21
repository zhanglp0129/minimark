package service

import (
	"gorm.io/gorm"
	"minimark/dao"
)

func PayMethodCreate(payMethod *dao.PayMethod) error {
	db := dao.GetDB()
	err := db.Create(payMethod).Error
	return err
}

func PayMethodList() ([]dao.PayMethod, error) {
	var payMethod []dao.PayMethod
	db := dao.GetDB()
	err := db.Find(&payMethod).Error
	return payMethod, err
}

func PayMethodFind(id int) (dao.PayMethod, error) {
	var payMethod dao.PayMethod
	db := dao.GetDB()
	err := db.Take(&payMethod, id).Error
	return payMethod, err
}

func PayMethodUpdate(payMethod *dao.PayMethod) error {
	db := dao.GetDB()
	err := db.Save(payMethod).Error
	return err
}

func PayMethodDelete(id int) error {
	db := dao.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		// 先将订单和进货的支付方式置为null
		err := tx.Model(&dao.Order{}).Where("pay_method_id = ?", id).Update("pay_method_id", nil).Error
		if err != nil {
			return err
		}
		err = tx.Model(&dao.Procurement{}).Where("pay_method_id = ?", id).Update("pay_method_id", nil).Error
		if err != nil {
			return err
		}

		// 再删除支付方式
		err = tx.Delete(&dao.PayMethod{ID: id}).Error
		if err != nil {
			return err
		}

		return nil
	})
	return err
}
