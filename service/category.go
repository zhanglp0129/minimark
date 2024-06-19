package service

import (
	"minimark/dao"
)

func CategoryCreate(category *dao.Category) error {
	db := dao.GetDB()
	err := db.Create(category).Error
	return err
}

func CategoryList() ([]dao.Category, error) {
	var categories []dao.Category
	db := dao.GetDB()
	err := db.Find(&categories).Error
	return categories, err
}

func CategoryFind(id int) (dao.Category, error) {
	var category dao.Category
	db := dao.GetDB()
	err := db.Take(&category, id).Error
	return category, err
}

func CategoryUpdate(category *dao.Category) error {
	db := dao.GetDB()
	err := db.Save(category).Error
	return err
}

func CategoryDelete(id int) error {
	db := dao.GetDB()
	err := db.Delete(&dao.Category{}, id).Error
	return err
}
