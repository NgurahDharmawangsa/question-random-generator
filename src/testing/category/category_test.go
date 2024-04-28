package category_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
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

func create(t *testing.T, name string, order uint) uint {
	Init()

	catData := models.Category{
		Name:  name,
		Order: order,
	}

	jsonData, err := json.Marshal(catData)
	if err != nil {
		t.Errorf("error encoding JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:3000/api/categories", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}

	var responseData struct {
		Data    models.Category `json:"data"`
		Message string          `json:"message"`
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		t.Errorf("error decoding JSON response: %v", err)
	}

	id := responseData.Data.ID
	return id
}

func rollback(t *testing.T, name string) {
	config.Mysql.DB.Unscoped().Where("name = ?", name).Delete(&models.Category{})
	//database.Instance.Where("name = ?", name).Delete(&entity.Category{})
}

func TestSuccessCreate(t *testing.T) {
	Init()

	catData := models.Category{
		Name:  "Cat Create",
		Order: 1,
	}

	jsonData, err := json.Marshal(catData)
	assert.Nil(t, err)

	req, err := http.NewRequest("POST", "http://127.0.0.1:3000/api/categories", bytes.NewBuffer(jsonData))
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	rollback(t, "Cat Create")
}

func TestSuccessGetAll(t *testing.T) {
	Init()

	resp, err := http.Get("http://127.0.0.1:3000/api/categories")
	assert.Nil(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	//fmt.Println(string(body))

	if len(body) == 0 {
		t.Errorf("empty response body")
	}
}

func TestSuccessGetByID(t *testing.T) {
	Init()

	id := create(t, "Category Get By ID", 2)

	fmt.Println(id)
	//time.Sleep(2 * time.Second)
	getByIDURL := fmt.Sprintf("http://127.0.0.1:3000/api/categories/by-id/%d", id)

	resp, err := http.Get(getByIDURL)
	if err != nil {
		t.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code when getting category by ID: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	if len(body) == 0 {
		t.Errorf("empty response body")
	}

	//if err != nil {
	//	t.Errorf("error reading response body: %v", err)
	//}
	//
	//fmt.Println("Response body:", string(body))

	rollback(t, "Category Get By ID")
}

func TestUpdateData(t *testing.T) {
	Init()

	id := create(t, "Category Sink", 5)
	catData := models.Category{
		Name:  "Cat Update",
		Order: 4,
	}

	jsonData, err := json.Marshal(catData)
	assert.Nil(t, err)

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://127.0.0.1:3000/api/categories/by-id/%d", id), bytes.NewBuffer(jsonData))
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	rollback(t, "Cat Update")
}

func TestDeleteData(t *testing.T) {
	Init()

	id := create(t, "Category Delete", 5)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://127.0.0.1:3000/api/categories/by-id/%d", id), nil)
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	rollback(t, "Category Delete")
}
