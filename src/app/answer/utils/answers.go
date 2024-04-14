package utils

import (
	"sekolahbeta/final-project/question-random-generator/src/app/answer/model"
	"sekolahbeta/final-project/question-random-generator/src/config"
	"time"
)

func InsertAnswerData(data model.Answer) (model.Answer, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Mysql.DB)

	return data, err
}

func GetAnswersList() ([]model.Answer, error) {
	var answer model.Answer
	return answer.GetAllAnswer(config.Mysql.DB)
}

func GetAnswersByID(id uint) (model.Answer, error) {
	answer := model.Answer{
		Model: model.Model{
			ID: id,
		},
	}
	return answer.GetByID(config.Mysql.DB)
}

func UpdateAnswersByID(data model.Answer, id uint) (model.Answer, error) {
	data.UpdatedAt = time.Now()
	err := data.UpdateOneByID(config.Mysql.DB, id)

	return data, err
}

func DeleteByID(id uint) error {
	answer := model.Answer{
		Model: model.Model{
			ID: id,
		},
	}
	return answer.DeleteByID(config.Mysql.DB)
}
