package utils

import (
	"fmt"
	"math/rand"
	"sekolahbeta/final-project/question-random-generator/src/app/models"
	"sekolahbeta/final-project/question-random-generator/src/config"
	"sort"
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
	seed := time.Now().UnixNano()
	randUtil := rand.New(rand.NewSource(seed))

	module := models.Module{
		Identifier: identifier,
	}
	question := models.Question{}

	a, err := module.GetQuestions(config.Mysql.DB)
	if err != nil {
		return a, err
	}
	
	var ids []any
	for _, id := range a.QuestionIds {
		ids = append(ids, id)
	}

	b, err := question.GetQuestionByIDS(config.Mysql.DB, ids)
	if err != nil {
		return a, err
	}

	// var orderTemp string 
	ansRand := []models.Answer{}
	for i, v := range b {
		randUtil.Shuffle(len(v.Answer), func(i, j int) {
			v.Answer[i], v.Answer[j] = v.Answer[j], v.Answer[i]
		})
		ansRand = append(ansRand, v.Answer...)

		b[i].Answer = ansRand
	}

	catGroup := make(map[uint][]models.Question)
	for _, v := range b {
		catGroup[v.Category.Order] = append(catGroup[v.Category.Order], v)
	}

	fmt.Println(catGroup)
	for _, group := range catGroup {
		randUtil.Shuffle(len(group), func(i, j int) {
			group[i], group[j] = group[j], group[i]
		})
	}

	var sortedCat []uint
	for k := range catGroup {
		sortedCat = append(sortedCat, k)
	}

	sort.Slice(sortedCat, func(i, j int) bool {
		return sortedCat[i] < sortedCat[j]
	})

	var qstTemp []models.Question
	for _, v := range sortedCat {
		group := catGroup[v]
		qstTemp = append(qstTemp, group...)
	}

	for k := range qstTemp {
		qstTemp[k].Category.Order = uint(k+1)
	}

	a.Question = qstTemp

	return a, nil
}
