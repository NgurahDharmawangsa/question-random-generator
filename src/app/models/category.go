package models

import "gorm.io/gorm"

type Category struct {
	Model
	Name  string `gorm:"not null" json:"name"`
	Order uint `gorm:"not null" json:"order"`
}
type Categories []Category

func (cat *Category) Create(db *gorm.DB) error {
	err := db.
		Model(Category{}).
		Create(&cat).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cat *Category) GetAll(db *gorm.DB) ([]Category, error) {
	res := []Category{}

	err := db.
		Model(Category{}).
		Find(&res).
		Error

	if err != nil {
		return []Category{}, err
	}

	return res, nil
}

func (cat *Category) GetByID(db *gorm.DB) (Category, error) {
	res := Category{}

	err := db.
		Model(Category{}).
		Where("id = ?", cat.Model.ID).
		Take(&res).
		Error

	if err != nil {
		return Category{}, err
	}

	return res, nil
}

func (cat *Category) UpdateOneByID(db *gorm.DB, id uint) error {
	err := db.
		Model(Category{}).
		Select("name", "order").
		Where("id = ?", id).
		Updates(map[string]any{
			"name":  cat.Name,
			"order": cat.Order,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cat *Category) DeleteByID(db *gorm.DB) error {
	err := db.
		Model(Category{}).
		Where("id = ?", cat.Model.ID).
		Delete(&cat).
		Error

	if err != nil {
		return err
	}

	return nil
}
