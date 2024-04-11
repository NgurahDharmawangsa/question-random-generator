package model

import (
	// "sekolahbeta/final-project/question-random-generator/src/app/question/model"

	"gorm.io/gorm"
)

type Answer struct {
	Model
	Option     string         `gorm:"not null" json:"option"`
	Answer     string         `gorm:"not null" json:"answer"`
	Score      int            `gorm:"not null"  json:"score"`
	QuestionId string         `gorm:"type:char(36);constraint:OnDelete:CASCADE;" json:"question_id"`
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
