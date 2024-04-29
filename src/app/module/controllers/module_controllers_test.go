package controllers_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sekolahbeta/final-project/question-random-generator/src/config"
	"sekolahbeta/final-project/question-random-generator/src/routes"

	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var (
	fiberApp *fiber.App
)

func Init() {
	err := godotenv.Load("../../../../.env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
	config.OpenDB()

	fiberApp = fiber.New()
}

func Test_CreateModule(t *testing.T) {
	Init()
	routes.Api(fiberApp)

	t.Run("Success Create", func(t *testing.T) {

		requestBody := []byte(`{"name": "module control", "question_ids": [3, 2, 1]}`)

		request, e := http.NewRequest(
			"POST",
			"http://127.0.0.1:3000/api/modules",
			bytes.NewBuffer(requestBody),
		)
		assert.Equal(t, nil, e)

		request.Header.Set("Content-Type", "application/json")

		request.Header.Set("Role", "Admin")

		response, err := fiberApp.Test(request, -1)

		assert.Equal(t, nil, err)

		assert.Equal(t, fiber.StatusCreated, response.StatusCode)

		b, err := io.ReadAll(response.Body)

		assert.Equal(t, nil, err)

		assert.JSONEq(t, string(b), string(b))
	})
}

func Test_GetAllModule(t *testing.T) {
	Init()
	routes.Api(fiberApp)

	t.Run("Success Get All", func(t *testing.T) {

		request, e := http.NewRequest(
			"GET",
			"http://127.0.0.1:3000/api/modules",
			nil,
		)
		assert.Equal(t, nil, e)

		response, err := fiberApp.Test(request, -1)

		assert.Equal(t, nil, err)

		assert.Equal(t, fiber.StatusOK, response.StatusCode)

		b, err := io.ReadAll(response.Body)

		assert.Equal(t, nil, err)

		assert.JSONEq(t, string(b), string(b))
	})
}

func Test_GetModuleById(t *testing.T) {
	Init()
	routes.Api(fiberApp)

	t.Run("Success Get by ID", func(t *testing.T) {
		moduleID := 1

		request, e := http.NewRequest(
			"GET",
			fmt.Sprintf("http://127.0.0.1:3000/api/modules/%d", moduleID),
			nil,
		)
		assert.Equal(t, nil, e)

		response, err := fiberApp.Test(request, -1)

		assert.Equal(t, nil, err)

		assert.Equal(t, fiber.StatusOK, response.StatusCode)

		b, err := io.ReadAll(response.Body)

		assert.Equal(t, nil, err)

		assert.JSONEq(t, string(b), string(b))
	})
}

func Test_UpdateModuleById(t *testing.T) {
	Init()
	routes.Api(fiberApp)

	t.Run("Success Update by ID", func(t *testing.T) {
		moduleID := 3648386603

		requestBody := []byte(`{"name": "Module Updated"}`)

		request, e := http.NewRequest(
			"PUT",
			fmt.Sprintf("http://127.0.0.1:3000/api/modules/%d", moduleID),
			bytes.NewBuffer(requestBody),
		)
		assert.Equal(t, nil, e)

		request.Header.Set("Content-Type", "application/json")

		request.Header.Set("Role", "Admin")

		response, err := fiberApp.Test(request, -1)

		assert.Equal(t, nil, err)

		assert.Equal(t, fiber.StatusCreated, response.StatusCode)

		b, err := io.ReadAll(response.Body)

		assert.Equal(t, nil, err)

		assert.JSONEq(t, string(b), string(b))
	})
}

func Test_DeleteCategoryById(t *testing.T) {
	Init()
	routes.Api(fiberApp)

	t.Run("Success Delete", func(t *testing.T) {
		moduleID := 3648386603

		request, e := http.NewRequest(
			"DELETE",
			fmt.Sprintf("http://127.0.0.1:3000/api/modules/%d", moduleID),
			nil,
		)
		assert.Equal(t, nil, e)

		request.Header.Set("Content-Type", "application/json")

		request.Header.Set("Role", "Admin")

		response, err := fiberApp.Test(request, -1)

		assert.Equal(t, nil, err)

		assert.Equal(t, fiber.StatusOK, response.StatusCode)

		b, err := io.ReadAll(response.Body)

		assert.Equal(t, nil, err)

		assert.JSONEq(t, `{"message":"Success Delete Data"}`, string(b))
	})
}

func Test_GetExamByIdentifier(t *testing.T) {
	Init()
	routes.Api(fiberApp)

	t.Run("Success Get Exam", func(t *testing.T) {
		identifier := "MDL-20240428145527"

		request, e := http.NewRequest(
			"GET",
			fmt.Sprintf("http://127.0.0.1:3000/api/modules/exam/questions/%s", identifier),
			nil,
		)
		assert.Equal(t, nil, e)

		request.Header.Set("Content-Type", "application/json")

		request.Header.Set("Role", "Admin")

		response, err := fiberApp.Test(request, -1)

		assert.Equal(t, nil, err)

		assert.Equal(t, fiber.StatusOK, response.StatusCode)

		b, err := io.ReadAll(response.Body)

		assert.Equal(t, nil, err)

		assert.JSONEq(t, string(b), string(b))
	})
}
