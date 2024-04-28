package answer_test

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

func create(t *testing.T, option string, answer string, score int, questionId string) uint {
	Init()

	ansData := models.Answer{
		Option:     option,
		Answer:     answer,
		Score:      score,
		QuestionId: questionId,
	}

	jsonData, err := json.Marshal(ansData)
	if err != nil {
		t.Errorf("error encoding JSON: %v", err)
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:3000/api/answers", bytes.NewBuffer(jsonData))
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
		Data    models.Answer `json:"data"`
		Message string        `json:"message"`
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		t.Errorf("error decoding JSON response: %v", err)
	}
	id := responseData.Data.ID
	return id
}

func rollback(t *testing.T, answer string) {
	config.Mysql.DB.Unscoped().Where("answer = ?", answer).Delete(&models.Answer{})
}

func TestSuccessCreate(t *testing.T) {
	Init()

	ansData := models.Answer{
		Option:     "A",
		Answer:     "Answer Create",
		Score:      10,
		QuestionId: "1",
	}

	jsonData, err := json.Marshal(ansData)
	assert.Nil(t, err)

	req, err := http.NewRequest("POST", "http://127.0.0.1:3000/api/answers", bytes.NewBuffer(jsonData))
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	rollback(t, "Answer Create")
}

func TestSuccessGetAll(t *testing.T) {
	Init()

	resp, err := http.Get("http://127.0.0.1:3000/api/answers")
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

	id := create(t, "B", "Answer Get By ID", 10, "1")

	getByIDURL := fmt.Sprintf("http://127.0.0.1:3000/api/answers/%d", id)

	resp, err := http.Get(getByIDURL)
	if err != nil {
		t.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code when getting answers by ID: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	if len(body) == 0 {
		t.Errorf("empty response body")
	}

	rollback(t, "Answer Get By ID")
}

func TestUpdateData(t *testing.T) {
	Init()

	id := create(t, "C", "Answer Before Update", 25, "1")
	ansData := models.Answer{
		Option:     "D",
		Answer:     "Answer Updated",
		Score:      15,
		QuestionId: "1",
	}

	jsonData, err := json.Marshal(ansData)
	assert.Nil(t, err)

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://127.0.0.1:3000/api/answers/%d", id), bytes.NewBuffer(jsonData))
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	rollback(t, "Answer Updated")
}

func TestDeleteData(t *testing.T) {
	Init()

	id := create(t, "E", "Answer Deleted", 1, "1")

	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://127.0.0.1:3000/api/answers/%d", id), nil)
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	rollback(t, "Answer Deleted")
}
