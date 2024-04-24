package utils

import (
	"sekolahbeta/final-project/question-random-generator/src/app/models"
	"sekolahbeta/final-project/question-random-generator/src/config"
	"time"
)

func InsertModuleData(data models.Module) (models.Module, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	err := data.Create(config.Mysql.DB)

	return data, err
}

func GetModulesList() ([]models.Module, error) {
	var answer models.Module
	return answer.GetAllModul(config.Mysql.DB)
}

func GetModulesByID(id uint) (models.Module, error) {
	module := models.Module{
		Model: models.Model{
			ID: id,
		},
	}
	return module.GetByID(config.Mysql.DB)
}

func UpdateModulesByID(data models.Module, id uint) (models.Module, error) {
	data.UpdatedAt = time.Now()
	err := data.UpdateOneByID(config.Mysql.DB, id)

	return data, err
}

func DeleteByID(id uint) error {
	module := models.Module{
		Model: models.Model{
			ID: id,
		},
	}
	return module.DeleteByID(config.Mysql.DB)
}

func GetQuestions(identifier string) (models.Module, error) {
	module := models.Module{
		Identifier: identifier,
	}
	return module.GetQuestions(config.Mysql.DB)
}
