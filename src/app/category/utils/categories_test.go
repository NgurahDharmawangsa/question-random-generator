package utils_test

import (
	"fmt"
	"sekolahbeta/final-project/question-random-generator/src/app/category/utils"
	"sekolahbeta/final-project/question-random-generator/src/app/models"
	"sekolahbeta/final-project/question-random-generator/src/config"

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

func TestInsertCategoryData_Success(t *testing.T) {
	Init()

	category := models.Category{
		Name:  "Test Utils Category",
		Order: 22,
	}

	_, err := utils.InsertCategoryData(category)
	if err != nil {
		t.Errorf("Error inserting category: %v", err)
	}

	config.Mysql.DB.Unscoped().Where("name = ?", category.Name).Delete(&category)
}

func TestInsertCategoryData_Failed(t *testing.T) {
	Init()

	category := models.Category{
		Name:  "",
		Order: 90,
	}

	_, err := utils.InsertCategoryData(category)

	if err != nil {
		t.Error("Expected an error but got none")
	}

	config.Mysql.DB.Unscoped().Where("name = ?", category.Name).Delete(&category)
}

func TestGetCategoriesList(t *testing.T) {
	Init()

	categories := []models.Category{
		{Name: "Category List", Order: 21},
	}

	for _, category := range categories {
		_, err := utils.InsertCategoryData(category)
		if err != nil {
			t.Fatalf("Failed to insert category: %v", err)
		}
	}

	_, err := utils.GetCategoriesList()

	if err != nil {
		t.Fatalf("Failed to get categories list: %v", err)
	}

	config.Mysql.DB.Unscoped().Where("name = ?", categories[0].Name).Delete(&categories)
}

func TestGetCategoryByID(t *testing.T) {
	Init()

	category := models.Category{
		Name:  "Test Category Get By ID",
		Order: 12,
	}

	insertedCategory, err := utils.InsertCategoryData(category)
	if err != nil {
		t.Fatalf("Failed to insert category: %v", err)
	}

	_, err = utils.GetCategoryByID(insertedCategory.ID)
	if err != nil {
		t.Fatalf("Failed to get category by ID: %v", err)
	}

	config.Mysql.DB.Unscoped().Where("name = ?", category.Name).Delete(&category)
}

func TestUpdateCategoriesByID(t *testing.T) {
	Init()

	category := models.Category{
		Name:  "Test Update Category",
		Order: 19,
	}

	insertedCategory, err := utils.InsertCategoryData(category)
	if err != nil {
		t.Fatalf("Failed to insert category: %v", err)
	}

	updatedName := "Updated Category"
	insertedCategory.Name = updatedName

	_, err = utils.UpdateCategoriesByID(insertedCategory, insertedCategory.ID)
	if err != nil {
		t.Fatalf("Failed to update category by ID: %v", err)
	}

	config.Mysql.DB.Unscoped().Where("name = ?", insertedCategory.Name).Delete(&category)
}

func TestDeleteByID(t *testing.T) {
	Init()

	category := models.Category{
		Name:  "Test Delete Category",
		Order: 99,
	}

	insertedCategory, err := utils.InsertCategoryData(category)
	if err != nil {
		t.Fatalf("Failed to insert category: %v", err)
	}

	err = utils.DeleteByID(insertedCategory.ID)
	if err != nil {
		t.Fatalf("Failed to delete category by ID: %v", err)
	}

	_, err = utils.GetCategoryByID(insertedCategory.ID)
	if err == nil {
		t.Fatalf("Expected category to be deleted, but it still exists")
	}

	config.Mysql.DB.Unscoped().Where("name = ?", category.Name).Delete(&category)
}
