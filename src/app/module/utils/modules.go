package utils

import (
	"sekolahbeta/final-project/question-random-generator/src/app/module/model"
	"sekolahbeta/final-project/question-random-generator/src/config"
	"time"
)

func InsertModuleData(data model.Module) (model.Module, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Mysql.DB)

	return data, err
}