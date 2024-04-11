package model

import (
	ans "sekolahbeta/final-project/question-random-generator/src/app/answer/model"
	cat "sekolahbeta/final-project/question-random-generator/src/app/category/model"

	"gorm.io/gorm"
)

type Question struct {
	Model
	Question   string       `gorm:"not null" json:"question"`
	CategoryId string       `gorm:"type:char(36);" json:"category_id"`
	Category   cat.Category `gorm:"constraint:OnDelete:CASCADE;"`
	Answer     []ans.Answer `gorm:"constraint:OnDelete:CASCADE;"`
}

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
		Preload("Answer").
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
