package utils

import (
	"sekolahbeta/final-project/question-random-generator/src/app/category/model"
	"sekolahbeta/final-project/question-random-generator/src/config"
	"time"
)

func InsertCategoryData(data model.Category) (model.Category, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Mysql.DB)

	return data, err
}

func GetCategoriesList() ([]model.Category, error) {
	var category model.Category
	return category.GetAll(config.Mysql.DB)
}

func GetCategoryByID(id uint) (model.Category, error) {
	category := model.Category{
		Model: model.Model{
			ID: id,
		},
	}
	return category.GetByID(config.Mysql.DB)
}

func UpdateCategoriesByID(data model.Category, id uint) (model.Category, error) {
	data.UpdatedAt = time.Now()
	err := data.UpdateOneByID(config.Mysql.DB, id)

	return data, err
}

func DeleteByID(id uint) error {
	category := model.Category{
		Model: model.Model{
			ID: id,
		},
	}
	return category.DeleteByID(config.Mysql.DB)
}
