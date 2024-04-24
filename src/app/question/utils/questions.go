package utils

import (
	"sekolahbeta/final-project/question-random-generator/src/app/models"
	"sekolahbeta/final-project/question-random-generator/src/config"
	"time"
)

func InsertQuestionData(data models.Question) (models.Question, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Mysql.DB)

	return data, err
}

func GetQuestionsList() ([]models.Question, error) {
	var category models.Question
	return category.GetAllQuestion(config.Mysql.DB)
}

func GetQuestionsByID(id uint) (models.Question, error) {
	question := models.Question{
		Model: models.Model{
			ID: id,
		},
	}
	return question.GetByID(config.Mysql.DB)
}

func UpdateQuestionsByID(data models.Question, id uint) (models.Question, error) {
	data.UpdatedAt = time.Now()
	err := data.UpdateOneByID(config.Mysql.DB, id)

	return data, err
}

func DeleteByID(id uint) error {
	question := models.Question{
		Model: models.Model{
			ID: id,
		},
	}
	return question.DeleteByID(config.Mysql.DB)
}
