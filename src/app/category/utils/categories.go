package utils

import (
	"sekolahbeta/final-project/question-random-generator/src/app/models"
	"sekolahbeta/final-project/question-random-generator/src/config"
	"time"
)

func InsertCategoryData(data models.Category) (models.Category, error) {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	err := data.Create(config.Mysql.DB)

	return data, err
}

func GetCategoriesList() ([]models.Category, error) {
	var category models.Category
	return category.GetAll(config.Mysql.DB)
}

func GetCategoryByID(id uint) (models.Category, error) {
	category := models.Category{
		Model: models.Model{
			ID: id,
		},
	}
	return category.GetByID(config.Mysql.DB)
}

func UpdateCategoriesByID(data models.Category, id uint) (models.Category, error) {
	data.UpdatedAt = time.Now()
	err := data.UpdateOneByID(config.Mysql.DB, id)

	return data, err
}

func DeleteByID(id uint) error {
	category := models.Category{
		Model: models.Model{
			ID: id,
		},
	}
	return category.DeleteByID(config.Mysql.DB)
}
