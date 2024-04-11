package model

import (
	"gorm.io/gorm"
	"sekolahbeta/final-project/question-random-generator/src/app/category/model"
)

type Question struct {
	Model
	Question   string         `gorm:"not null" json:"question"`
	CategoryId string         `gorm:"type:char(36);" json:"category_id"`
	Category   model.Category `gorm:"constraint:OnDelete:CASCADE;"`
}
type Questions []Question

func (que *Question) Create(db *gorm.DB) error {
	err := db.
		Model(Question{}).
		Create(&que).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (que *Question) GetAllQuestion(db *gorm.DB) ([]Question, error) {
	res := []Question{}

	err := db.
		Model(Question{}).
		Preload("Category").
		Find(&res).
		Error

	if err != nil {
		return []Question{}, err
	}

	return res, nil
}

func (que *Question) GetByID(db *gorm.DB) (Question, error) {
	res := Question{}

	err := db.
		Model(Question{}).
		Preload("Category").
		Where("id = ?", que.Model.ID).
		Take(&res).
		Error

	if err != nil {
		return Question{}, err
	}

	return res, nil
}

func (que *Question) UpdateOneByID(db *gorm.DB, id uint) error {
	err := db.
		Model(Question{}).
		Preload("Category").
		Select("question", "category_id").
		Where("id = ?", id).
		Updates(map[string]any{
			"question":    que.Question,
			"category_id": que.CategoryId,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (que *Question) DeleteByID(db *gorm.DB) error {
	err := db.
		Model(Question{}).
		Where("id = ?", que.Model.ID).
		Delete(&que).
		Error

	if err != nil {
		return err
	}

	return nil
}
