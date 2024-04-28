package module_test

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
	"time"
)

func Init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
	config.OpenDB()
}

func create(t *testing.T, name string, ids []int64) uint {
	Init()

	ansData := models.Module{
		Name:        name,
		QuestionIds: ids,
	}

	jsonData, err := json.Marshal(ansData)
	if err != nil {
		t.Errorf("error encoding JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:3000/api/modules", bytes.NewBuffer(jsonData))
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
		Data    models.Module `json:"data"`
		Message string        `json:"message"`
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		t.Errorf("error decoding JSON response: %v", err)
	}
	id := responseData.Data.ID
	return id
}

func rollback(t *testing.T, name string) {
	config.Mysql.DB.Unscoped().Where("name = ?", name).Delete(&models.Module{})
}

func TestSuccessCreate(t *testing.T) {
	Init()

	modData := models.Module{
		Identifier:  "MDL-" + time.Now().Format("20060102150405"),
		Name:        "Module Create",
		QuestionIds: []int64{1, 2, 3},
	}

	jsonData, err := json.Marshal(modData)
	assert.Nil(t, err)

	req, err := http.NewRequest("POST", "http://127.0.0.1:3000/api/modules", bytes.NewBuffer(jsonData))
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	rollback(t, "Module Create")
}

func TestSuccessGetAll(t *testing.T) {
	Init()

	resp, err := http.Get("http://127.0.0.1:3000/api/modules")
	assert.Nil(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	if len(body) == 0 {
		t.Errorf("empty response body")
	}
}

func TestSuccessGetByID(t *testing.T) {
	Init()

	id := create(t, "Module Create", []int64{1, 2, 3})
	time.Sleep(2 * time.Second)
	fmt.Println(id)

	getByIDURL := fmt.Sprintf("http://127.0.0.1:3000/api/modules/%d", id)

	resp, err := http.Get(getByIDURL)
	if err != nil {
		t.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code when getting modules by ID: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	if len(body) == 0 {
		t.Errorf("empty response body")
	}

	rollback(t, "Module Get By ID")
}

func TestUpdateData(t *testing.T) {
	Init()

	id := create(t, "Module Before Update", []int64{1, 2})
	modData := models.Module{
		Identifier: "MDL-22",
		Name:       "Module Updated",
	}

	jsonData, err := json.Marshal(modData)
	assert.Nil(t, err)

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://127.0.0.1:3000/api/modules/%d", id), bytes.NewBuffer(jsonData))
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	rollback(t, "Module Updated")
}

func TestDeleteData(t *testing.T) {
	Init()

	id := create(t, "Module Deleted", []int64{1, 2})

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://127.0.0.1:3000/api/modules/%d", id), nil)
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	rollback(t, "Module Deleted")
}

func TestSuccessGetByIdentifier(t *testing.T) {
	Init()

	identifier := "MDL-20240428145900"

	getByIDURL := fmt.Sprintf("http://127.0.0.1:3000/api/modules/exam/questions/%s", identifier)

	resp, err := http.Get(getByIDURL)
	if err != nil {
		t.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code when getting modules by ID: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	if len(body) == 0 {
		t.Errorf("empty response body")
	}
}
