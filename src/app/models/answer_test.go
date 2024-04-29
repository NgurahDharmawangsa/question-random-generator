package models_test

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"sekolahbeta/final-project/question-random-generator/src/app/models"
	"sekolahbeta/final-project/question-random-generator/src/config"
	"testing"
)

func Init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
	config.OpenDB()
}

func TestCreateAnswerSuccess(t *testing.T) {
	Init()

	ans := models.Answer{
		Option:     "A",
		Answer:     "Answer A",
		Score:      10,
		QuestionId: "1",
	}

	err := ans.Create(config.Mysql.DB)
	assert.NoError(t, err)

	config.Mysql.DB.Unscoped().Delete(&ans)
}

func TestGetByIdAnswerSuccess(t *testing.T) {
	Init()

	ans := models.Answer{
		Option:     "A",
		Answer:     "Answer A",
		Score:      10,
		QuestionId: "1",
	}

	err := ans.Create(config.Mysql.DB)
	assert.Nil(t, err)

	ans = models.Answer{
		Model: models.Model{
			ID: ans.ID,
		},
	}

	_, err = ans.GetByID(config.Mysql.DB)
	assert.Nil(t, err)

	config.Mysql.DB.Unscoped().Delete(&ans)
}

func TestGetAllAnswerSuccess(t *testing.T) {
	Init()

	ans := models.Answer{
		Option:     "A",
		Answer:     "Answer A",
		Score:      10,
		QuestionId: "1",
	}

	err := ans.Create(config.Mysql.DB)
	assert.Nil(t, err)

	res, err := ans.GetAllAnswer(config.Mysql.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(res), 1)

	config.Mysql.DB.Unscoped().Delete(&ans)
}

func TestUpdateByIDAnswer(t *testing.T) {
	Init()

	ans := models.Answer{
		Option:     "A",
		Answer:     "Answer A",
		Score:      10,
		QuestionId: "1",
	}

	err := ans.Create(config.Mysql.DB)
	assert.Nil(t, err)

	ans.Answer = "Answer Updated"

	err = ans.UpdateOneByID(config.Mysql.DB, ans.ID)
	assert.Nil(t, err)

	config.Mysql.DB.Unscoped().Delete(&ans)
}

func TestDeleteByIDAnswer(t *testing.T) {
	Init()

	ans := models.Answer{
		Option:     "A",
		Answer:     "Answer A",
		Score:      10,
		QuestionId: "1",
	}

	err := ans.Create(config.Mysql.DB)
	assert.Nil(t, err)

	ans = models.Answer{
		Model: models.Model{
			ID: ans.ID,
		},
	}

	err = ans.DeleteByID(config.Mysql.DB)
	assert.Nil(t, err)

	config.Mysql.DB.Unscoped().Delete(&ans)
}
