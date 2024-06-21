package service

import (
	"errors"
	"minimark/dao"
	"minimark/pojo"
)

func GoodsCreate(goods *dao.Goods) error {
	db := dao.GetDB()
	err := db.Create(goods).Error
	return err
}

func GoodsPage(dto pojo.GoodsPageDTO) ([]dao.Goods, error) {
	db := dao.GetDB()
	offset := (dto.PageNum - 1) * dto.PageSize
	var goods []dao.Goods
	tx := db.Preload("Category").Order("id desc").Offset(offset).Limit(dto.PageSize)
	if dto.GoodsName != nil {
		tx = tx.Where("name like ?", "%"+*dto.GoodsName+"%")
	}
	if dto.CategoryID != nil {
		tx = tx.Where("category_id = ?", *dto.CategoryID)
	}
	err := tx.Find(&goods).Error
	return goods, err
}

func GoodsFind(ids []int) ([]dao.Goods, error) {
	db := dao.GetDB()
	var goods []dao.Goods
	err := db.Preload("Category").Find(&goods, ids).Error
	if goods == nil || len(ids) != len(goods) {
		return nil, errors.New("商品id不存在")
	}
	return goods, err
}

func GoodsUpdate(goods *dao.Goods) error {
	db := dao.GetDB()
	err := db.Save(goods).Error
	return err
}

func GoodsDelete(id int) error {
	db := dao.GetDB()
	err := db.Delete(&dao.Goods{}, id).Error
	return err
}
