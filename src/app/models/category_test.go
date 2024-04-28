package models_test

import (
	"fmt"
	"sekolahbeta/final-project/question-random-generator/src/app/models"
	"sekolahbeta/final-project/question-random-generator/src/config"

	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
	config.OpenDB()
}

func TestCreateCateogrySuccess(t *testing.T) {
	Init()

	cat := models.Category{
		Name:  "Kategori Baru",
		Order: 10,
	}

	err := cat.Create(config.Mysql.DB)
	assert.NoError(t, err)

	config.Mysql.DB.Unscoped().Delete(&cat)
}

func TestGetByIdCateogrySuccess(t *testing.T) {
	Init()

	cat := models.Category{
		Name:  "Kategori Get By ID",
		Order: 11,
	}

	err := cat.Create(config.Mysql.DB)
	assert.Nil(t, err)

	cat = models.Category{
		Model: models.Model{
			ID: cat.ID,
		},
	}

	_, err = cat.GetByID(config.Mysql.DB)
	assert.Nil(t, err)

	config.Mysql.DB.Unscoped().Delete(&cat)
}

func TestGetAllCategorySuccess(t *testing.T) {
	Init()

	cat := models.Category{
		Name:  "Kategori Get All",
		Order: 12,
	}

	err := cat.Create(config.Mysql.DB)
	assert.Nil(t, err)

	res, err := cat.GetAll(config.Mysql.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(res), 1)

	config.Mysql.DB.Unscoped().Delete(&cat)
}

func TestUpdateByIDCategory(t *testing.T) {
	Init()

	cat := models.Category{
		Name:  "Category Update",
		Order: 13,
	}

	err := cat.Create(config.Mysql.DB)
	assert.Nil(t, err)

	cat.Name = "Category New Update"

	err = cat.UpdateOneByID(config.Mysql.DB, cat.ID)
	assert.Nil(t, err)

	config.Mysql.DB.Unscoped().Delete(&cat)
}

func TestDeleteByID(t *testing.T) {
	Init()

	cat := models.Category{
		Name: "Category Deleted",
		Order: 14,
	}

	err := cat.Create(config.Mysql.DB)
	assert.Nil(t, err)

	cat = models.Category{
		Model: models.Model{
			ID: cat.ID,
		},
	}

	err = cat.DeleteByID(config.Mysql.DB)
	assert.Nil(t, err)

	config.Mysql.DB.Unscoped().Delete(&cat)
}
