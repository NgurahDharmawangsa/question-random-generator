package utils

import (
	"sekolahbeta/final-project/question-random-generator/src/app/question/model"
	"sekolahbeta/final-project/question-random-generator/src/config"
	"time"
)

func InsertQuestionData(data model.Question) (model.Question, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Mysql.DB)

	return data, err
}

func GetQuestionsList() ([]model.Question, error) {
	var category model.Question
	return category.GetAllQuestion(config.Mysql.DB)
}

func GetQuestionsByID(id uint) (model.Question, error) {
	question := model.Question{
		Model: model.Model{
			ID: id,
		},
	}
	return question.GetByID(config.Mysql.DB)
}

func UpdateQuestionsByID(data model.Question, id uint) (model.Question, error) {
	data.UpdatedAt = time.Now()
	err := data.UpdateOneByID(config.Mysql.DB, id)

	return data, err
}

func DeleteByID(id uint) error {
	question := model.Question{
		Model: model.Model{
			ID: id,
		},
	}
	return question.DeleteByID(config.Mysql.DB)
}
