package model

import (
	// "sekolahbeta/final-project/question-random-generator/src/app/question/model"

	"gorm.io/gorm"
)

type Answer struct {
	Model
	Option     string `gorm:"not null" json:"option"`
	Answer     string `gorm:"not null" json:"answer"`
	Score      int    `gorm:"not null"  json:"score"`
	QuestionId string `gorm:"type:char(36);constraint:OnDelete:CASCADE;" json:"question_id"`
	// Question   model.Question `gorm:"constraint:OnDelete:CASCADE;"`
}

func (ans *Answer) Create(db *gorm.DB) error {
	err := db.
		Model(Answer{}).
		Preload("Question").
		Create(&ans).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (ans *Answer) GetAllAnswer(db *gorm.DB) ([]Answer, error) {
	res := []Answer{}

	err := db.
		Model(Answer{}).
		// Preload("Question.Category").
		Find(&res).
		Error

	if err != nil {
		return []Answer{}, err
	}

	return res, nil
}

func (ans *Answer) GetByID(db *gorm.DB) (Answer, error) {
	res := Answer{}

	err := db.
		Model(Answer{}).
		Where("id = ?", ans.Model.ID).
		Take(&res).
		Error

	if err != nil {
		return Answer{}, err
	}

	return res, nil
}

func (ans *Answer) UpdateOneByID(db *gorm.DB, id uint) error {
	err := db.
		Model(Answer{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"option":      ans.Option,
			"answer":      ans.Answer,
			"score":       ans.Score,
			"question_id": ans.QuestionId,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (ans *Answer) DeleteByID(db *gorm.DB) error {
	err := db.
		Model(Answer{}).
		Where("id = ?", ans.Model.ID).
		Delete(&ans).
		Error

	if err != nil {
		return err
	}

	return nil
}
