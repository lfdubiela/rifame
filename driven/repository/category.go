package repository

import (
	"github.com/jinzhu/gorm"
	"rifame/domain"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func (c *CategoryRepository) Save(category domain.Category) error {
	if err := c.DB.Create(&category).Error; err != nil {
		return err
	}

	return nil
}

func (c *CategoryRepository) FindAll() ([]*domain.Category, error) {
	var categories []*domain.Category

	err := c.DB.Order("name desc").Find(&categories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return categories, nil
}

//func (c *CategoryRepository) FindAll(pageNum int, pageSize int, maps interface{}) ([]*domain.Category, error) {
//	var categories []*domain.Category
//
//	err := c.DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&categories).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return nil, err
//	}
//
//	return categories, nil
//}

func (c *CategoryRepository) Get(id int) (*domain.Category, error) {
	var category domain.Category

	err := c.DB.Where("id = ? AND deleted_at = ? ", id, 0).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &category, nil
}

func (c *CategoryRepository) Delete(id int) error {
	if err := c.DB.Where("id = ?", id).Delete(&domain.Category{}).Error; err != nil {
		return err
	}

	return nil
}

func (c *CategoryRepository) Edit(id int, data interface{}) error {
	if err := c.DB.Model(&domain.Category{}).Where("id = ? AND deleted_at = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
