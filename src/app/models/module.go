package models

import (
	"math/rand"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Module struct {
	Model
	Identifier  string        `gorm:"not null" json:"identifier"`
	Name        string        `gorm:"not null" json:"name"`
	QuestionIds pq.Int64Array `gorm:"type:string" json:"question_ids"`
	Question    []Question    `gorm:"-" json:"question"`
}

func (mod *Module) Create(db *gorm.DB) error {
	err := db.
		Model(Module{}).
		Create(&mod).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (mod *Module) GetAllModul(db *gorm.DB) ([]Module, error) {
	res := []Module{}

	err := db.
		Model(Module{}).
		Find(&res).
		Error

	if err != nil {
		return []Module{}, err
	}

	return res, nil
}

func (mod *Module) GetByID(db *gorm.DB) (Module, error) {
	res := Module{}
	var questions []Question

	err := db.
		Model(Module{}).
		Where("id = ?", mod.Model.ID).
		Take(&res).
		Error

	if err != nil {
		return Module{}, err
	}

	var ids []interface{}
	for _, id := range res.QuestionIds {
		ids = append(ids, id)
	}

	err = db.
		Model(&Question{}).
		Where("id IN ?", ids).Preload("Answer").
		Find(&questions).
		Error
	if err != nil {
		return Module{}, err
	}

	res.Question = questions

	return res, nil
}

func (mod *Module) UpdateOneByID(db *gorm.DB, id uint) error {
	err := db.
		Model(Module{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"identifier":   mod.Identifier,
			"name":         mod.Name,
			"question_ids": mod.QuestionIds,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (mod *Module) DeleteByID(db *gorm.DB) error {
	err := db.
		Model(Module{}).
		Where("id = ?", mod.Model.ID).
		Delete(&mod).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (mod *Module) GetQuestions(db *gorm.DB) (Module, error) {
	res := Module{}
	questions := []Question{}

	err := db.
		Model(Module{}).
		Where("identifier = ?", mod.Identifier).
		Take(&res).
		Error

	if err != nil {
		return Module{}, err
	}

	var ids []interface{}
	for _, id := range res.QuestionIds {
		ids = append(ids, id)
	}

	err = db.
		Model(&Question{}).
		Where("id IN ?", ids).Preload("Category").Preload("Answer").
		Find(&questions).
		Error
	if err != nil {
		return Module{}, err
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	res.Question = questions

	return res, nil
}
