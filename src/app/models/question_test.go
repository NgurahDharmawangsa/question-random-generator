package models_test

import (
	"github.com/stretchr/testify/assert"
	"sekolahbeta/final-project/question-random-generator/src/app/models"
	"sekolahbeta/final-project/question-random-generator/src/config"
	"testing"
)

func TestCreateQuestionSuccess(t *testing.T) {
	Init()

	qst := models.Question{
		Question:   "Question A",
		CategoryId: "1",
	}

	err := qst.Create(config.Mysql.DB)
	assert.NoError(t, err)

	config.Mysql.DB.Unscoped().Delete(&qst)
}

func TestGetByIdQuestionSuccess(t *testing.T) {
	Init()

	qst := models.Question{
		Question:   "Question B",
		CategoryId: "1",
	}

	err := qst.Create(config.Mysql.DB)
	assert.Nil(t, err)

	qst = models.Question{
		Model: models.Model{
			ID: qst.ID,
		},
	}

	_, err = qst.GetByID(config.Mysql.DB)
	assert.Nil(t, err)

	config.Mysql.DB.Unscoped().Delete(&qst)
}

func TestGetAllQuestionSuccess(t *testing.T) {
	Init()

	qst := models.Question{
		Question:   "Question C",
		CategoryId: "1",
	}

	err := qst.Create(config.Mysql.DB)
	assert.Nil(t, err)

	res, err := qst.GetAllQuestion(config.Mysql.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(res), 1)

	config.Mysql.DB.Unscoped().Delete(&qst)
}

func TestUpdateByIDQuestion(t *testing.T) {
	Init()

	qst := models.Question{
		Question:   "Question D",
		CategoryId: "1",
	}

	err := qst.Create(config.Mysql.DB)
	assert.Nil(t, err)

	qst.Question = "Question Updated"

	err = qst.UpdateOneByID(config.Mysql.DB, qst.ID)
	assert.Nil(t, err)

	config.Mysql.DB.Unscoped().Delete(&qst)
}

func TestDeleteQuestionByID(t *testing.T) {
	Init()

	qst := models.Question{
		Question:   "Question Delete",
		CategoryId: "1",
	}

	err := qst.Create(config.Mysql.DB)
	assert.Nil(t, err)

	qst = models.Question{
		Model: models.Model{
			ID: qst.ID,
		},
	}

	err = qst.DeleteByID(config.Mysql.DB)
	assert.Nil(t, err)

	config.Mysql.DB.Unscoped().Delete(&qst)
}
