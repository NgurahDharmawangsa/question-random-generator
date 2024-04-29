package utils_test

import (
	"fmt"
	"sekolahbeta/final-project/question-random-generator/src/app/models"
	"sekolahbeta/final-project/question-random-generator/src/app/module/utils"
	"sekolahbeta/final-project/question-random-generator/src/config"
	"time"

	"testing"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
	config.OpenDB()
}

func TestInsertModuleData_Success(t *testing.T) {
	Init()

	module := models.Module{
		Identifier:  "MDL-" + time.Now().Format("20060102150405"),
		Name:        "Module 11",
		QuestionIds: []int64{4, 1, 2},
	}

	_, err := utils.InsertModuleData(module)
	if err != nil {
		t.Errorf("Error inserting module: %v", err)
	}

	config.Mysql.DB.Unscoped().Where("name = ?", module.Name).Delete(&module)
}

func TestInsertModuleData_Failed(t *testing.T) {
	Init()

	module := models.Module{
		Identifier:  "MDL-" + time.Now().Format("20060102150405"),
		Name:        "Module 12",
		QuestionIds: []int64{4, 1, 2},
	}

	_, err := utils.InsertModuleData(module)

	if err != nil {
		t.Error("Expected an error but got none")
	}

	config.Mysql.DB.Unscoped().Where("name = ?", module.Name).Delete(&module)
}

func TestGetModulesList(t *testing.T) {
	Init()

	modules := []models.Module{
		{
			Identifier:  "MDL-" + time.Now().Format("20060102150405"),
			Name:        "Module 13",
			QuestionIds: []int64{4, 1, 2},
		},
	}

	for _, module := range modules {
		_, err := utils.InsertModuleData(module)
		if err != nil {
			t.Fatalf("Failed to insert module: %v", err)
		}
	}

	_, err := utils.GetModulesList()

	if err != nil {
		t.Fatalf("Failed to get modules list: %v", err)
	}

	config.Mysql.DB.Unscoped().Where("name = ?", modules[0].Name).Delete(&modules)
}

func TestGetModuleByID(t *testing.T) {
	Init()

	module := models.Module{
		Identifier:  "MDL-" + time.Now().Format("20060102150405"),
		Name:        "Module 14",
		QuestionIds: []int64{4, 1, 2},
	}

	insertedModule, err := utils.InsertModuleData(module)
	if err != nil {
		t.Fatalf("Failed to insert module: %v", err)
	}

	_, err = utils.GetModulesByID(insertedModule.ID)
	if err != nil {
		t.Fatalf("Failed to get module by ID: %v", err)
	}

	config.Mysql.DB.Unscoped().Where("name = ?", module.Name).Delete(&module)
}

func TestUpdateModulesByID(t *testing.T) {
	Init()

	module := models.Module{
		Identifier:  "MDL-" + time.Now().Format("20060102150405"),
		Name:        "Module 15",
		QuestionIds: []int64{4, 5, 1},
	}

	insertedModule, err := utils.InsertModuleData(module)
	if err != nil {
		t.Fatalf("Failed to insert module: %v", err)
	}

	updatedName := "Updated Module"
	insertedModule.Name = updatedName

	_, err = utils.UpdateModulesByID(insertedModule, insertedModule.ID)
	if err != nil {
		t.Fatalf("Failed to update module by ID: %v", err)
	}

	config.Mysql.DB.Unscoped().Where("name = ?", insertedModule.Name).Delete(&module)
}

func TestDeleteByID(t *testing.T) {
	Init()

	module := models.Module{
		Identifier:  "MDL-" + time.Now().Format("20060102150405"),
		Name:        "Module 16",
		QuestionIds: []int64{4, 5, 1},
	}

	insertedModule, err := utils.InsertModuleData(module)
	if err != nil {
		t.Fatalf("Failed to insert module: %v", err)
	}

	err = utils.DeleteByID(insertedModule.ID)
	if err != nil {
		t.Fatalf("Failed to delete module by ID: %v", err)
	}

	_, err = utils.GetModulesByID(insertedModule.ID)
	if err == nil {
		t.Fatalf("Expected module to be deleted, but it still exists")
	}

	config.Mysql.DB.Unscoped().Where("name = ?", module.Name).Delete(&module)
}

func TestGetExamByIdentifier(t *testing.T) {
	Init()

	module := models.Module{
		Identifier:  "MDL-" + time.Now().Format("20060102150405"),
		Name:        "Module Identifier",
		QuestionIds: []int64{1, 2, 3, 4, 5},
	}

	insertedModule, err := utils.InsertModuleData(module)
	if err != nil {
		t.Fatalf("Failed to insert module: %v", err)
	}

	_, err = utils.GetQuestions(insertedModule.Identifier)
	if err != nil {
		t.Fatalf("Failed to get exam by identifier: %v", err)
	}

	config.Mysql.DB.Unscoped().Where("name = ?", module.Name).Delete(&module)
}
