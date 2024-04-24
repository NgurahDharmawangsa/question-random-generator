package model

import (
	"sekolahbeta/final-project/question-random-generator/src/app/category/model"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Module struct {
	Model
	Identifier  string           `gorm:"not null" json:"identifier"`
	Name        string           `gorm:"not null" json:"name"`
	QuestionIds pq.Int64Array    `gorm:"type:integer[]" json:"question_ids"`
	Question    []model.Category `gorm:"constraint:OnDelete:CASCADE;" json:"questions"`
}

func (mod *Module) Create(db *gorm.DB) error {
	for _, question := range mod.Question {
		mod.QuestionIds = append(mod.QuestionIds, int64(question.ID))
	}

	err := db.
		Model(Module{}).
		Preload("Question").
		Create(&mod).
		Error

	if err != nil {
		return err
	}

	return nil
}
