package question_test

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

func create(t *testing.T, question string, categoryId string) uint {
	Init()

	qstData := models.Question{
		Question:   question,
		CategoryId: categoryId,
	}

	jsonData, err := json.Marshal(qstData)
	if err != nil {
		t.Errorf("error encoding JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:3000/api/questions", bytes.NewBuffer(jsonData))
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
		Data    models.Question `json:"data"`
		Message string          `json:"message"`
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		t.Errorf("error decoding JSON response: %v", err)
	}

	id := responseData.Data.ID
	return id
}

func rollback(t *testing.T, question string) {
	config.Mysql.DB.Unscoped().Where("question = ?", question).Delete(&models.Question{})
}

func TestSuccessCreate(t *testing.T) {
	Init()

	qstData := models.Question{
		Question:   "Question 1",
		CategoryId: "1",
	}

	jsonData, err := json.Marshal(qstData)
	assert.Nil(t, err)

	req, err := http.NewRequest("POST", "http://127.0.0.1:3000/api/questions", bytes.NewBuffer(jsonData))
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	rollback(t, "Question 1")
}

func TestSuccessGetAll(t *testing.T) {
	Init()

	resp, err := http.Get("http://127.0.0.1:3000/api/questions")
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

	id := create(t, "Question Get By ID", "1")

	getByIDURL := fmt.Sprintf("http://127.0.0.1:3000/api/questions/by-id/%d", id)

	resp, err := http.Get(getByIDURL)
	if err != nil {
		t.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code when getting question by ID: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	if len(body) == 0 {
		t.Errorf("empty response body")
	}

	rollback(t, "Question Get By ID")
}

func TestUpdateData(t *testing.T) {
	Init()

	id := create(t, "Question Before Update", "1")
	qstData := models.Question{
		Question:   "Question Updated",
		CategoryId: "1",
	}

	jsonData, err := json.Marshal(qstData)
	assert.Nil(t, err)

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://127.0.0.1:3000/api/questions/by-id/%d", id), bytes.NewBuffer(jsonData))
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	rollback(t, "Question Updated")
}

func TestDeleteData(t *testing.T) {
	Init()

	id := create(t, "Question Delete", "1")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://127.0.0.1:3000/api/questions/by-id/%d", id), nil)
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	rollback(t, "Question Delete")
}
