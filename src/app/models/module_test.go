package models_test

import (
	"sekolahbeta/final-project/question-random-generator/src/app/models"
	"sekolahbeta/final-project/question-random-generator/src/config"
	"time"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateModuleSuccess(t *testing.T) {
	Init()

	mod := models.Module{
		Identifier:  "MDL-" + time.Now().Format("20060102150405"),
		Name:        "Module 1",
		QuestionIds: []int64{1, 2, 3},
	}

	err := mod.Create(config.Mysql.DB)
	assert.NoError(t, err)

	config.Mysql.DB.Unscoped().Delete(&mod)
}

func TestGetByIdModuleSuccess(t *testing.T) {
	Init()

	mod := models.Module{
		Identifier:  "MDL-" + time.Now().Format("20060102150405"),
		Name:        "Module 2",
		QuestionIds: []int64{1, 2, 3},
	}

	err := mod.Create(config.Mysql.DB)
	assert.Nil(t, err)

	mod = models.Module{
		Model: models.Model{
			ID: mod.ID,
		},
	}

	_, err = mod.GetByID(config.Mysql.DB)
	assert.Nil(t, err)

	config.Mysql.DB.Unscoped().Delete(&mod)
}

func TestGetAllModuleSuccess(t *testing.T) {
	Init()

	mod := models.Module{
		Identifier:  "MDL-" + time.Now().Format("20060102150405"),
		Name:        "Module 3",
		QuestionIds: []int64{1, 2, 3},
	}

	err := mod.Create(config.Mysql.DB)
	assert.Nil(t, err)

	res, err := mod.GetAllModul(config.Mysql.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(res), 1)

	config.Mysql.DB.Unscoped().Delete(&mod)
}

func TestUpdateByIDModule(t *testing.T) {
	Init()

	mod := models.Module{
		Identifier:  "MDL-" + time.Now().Format("20060102150405"),
		Name:        "Module 4",
		QuestionIds: []int64{1, 2, 3},
	}

	err := mod.Create(config.Mysql.DB)
	assert.Nil(t, err)

	mod.Name = "Module New Update"

	err = mod.UpdateOneByID(config.Mysql.DB, mod.ID)
	assert.Nil(t, err)

	config.Mysql.DB.Unscoped().Delete(&mod)
}

func TestDeleteModuleByID(t *testing.T) {
	Init()

	mod := models.Module{
		Identifier:  "MDL-" + time.Now().Format("20060102150405"),
		Name:        "Module 5",
		QuestionIds: []int64{4, 1, 2},
	}

	err := mod.Create(config.Mysql.DB)
	assert.Nil(t, err)

	mod = models.Module{
		Model: models.Model{
			ID: mod.ID,
		},
	}

	err = mod.DeleteByID(config.Mysql.DB)
	assert.Nil(t, err)

	config.Mysql.DB.Unscoped().Delete(&mod)
}
