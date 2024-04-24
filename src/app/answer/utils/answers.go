package utils

import (
	"sekolahbeta/final-project/question-random-generator/src/app/models"
	"sekolahbeta/final-project/question-random-generator/src/config"
	"time"
)

func InsertAnswerData(data models.Answer) (models.Answer, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Mysql.DB)

	return data, err
}

func GetAnswersList() ([]models.Answer, error) {
	var answer models.Answer
	return answer.GetAllAnswer(config.Mysql.DB)
}

func GetAnswersByID(id uint) (models.Answer, error) {
	answer := models.Answer{
		Model: models.Model{
			ID: id,
		},
	}
	return answer.GetByID(config.Mysql.DB)
}

func UpdateAnswersByID(data models.Answer, id uint) (models.Answer, error) {
	data.UpdatedAt = time.Now()
	err := data.UpdateOneByID(config.Mysql.DB, id)

	return data, err
}

func DeleteByID(id uint) error {
	answer := models.Answer{
		Model: models.Model{
			ID: id,
		},
	}
	return answer.DeleteByID(config.Mysql.DB)
}
