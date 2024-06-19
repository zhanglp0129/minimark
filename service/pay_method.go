package service

import "minimark/dao"

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
	err := db.Delete(&dao.PayMethod{}, id).Error
	return err
}
