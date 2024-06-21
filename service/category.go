package service

import (
	"gorm.io/gorm"
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
	err := db.Order("id desc").Find(&categories).Error
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
	err := db.Transaction(func(tx *gorm.DB) error {
		// 先将商品的分类置为null
		err := tx.Model(&dao.Goods{}).Where("category_id = ?", id).Update("category_id", nil).Error
		if err != nil {
			return err
		}

		// 再删除商品分类
		err = tx.Delete(&dao.Category{ID: id}).Error
		if err != nil {
			return err
		}

		return nil
	})
	return err
}
