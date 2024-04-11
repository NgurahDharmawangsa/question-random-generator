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
