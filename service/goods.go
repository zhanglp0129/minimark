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

func GoodsPage(dto pojo.GoodsPageDTO) ([]dao.Goods, int64, error) {
	db := dao.GetDB()
	var goods []dao.Goods
	// 获取查询条件
	tx := db.Model(&dao.Goods{})
	if dto.GoodsName != nil {
		tx = tx.Where("name like ?", "%"+*dto.GoodsName+"%")
	}
	if dto.CategoryID != nil {
		tx = tx.Where("category_id = ?", *dto.CategoryID)
	}
	// 获取总结果数
	var total int64
	err := tx.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 正式分页查询
	offset := (dto.PageNum - 1) * dto.PageSize
	tx = tx.Preload("Category").Order("id desc").Offset(offset).Limit(dto.PageSize)
	err = tx.Find(&goods).Error
	if err != nil {
		return nil, 0, err
	}

	return goods, total, err
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
