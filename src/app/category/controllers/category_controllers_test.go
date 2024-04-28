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

func Test_CreateCategory(t *testing.T) {
	Init()
	routes.Api(fiberApp)

	t.Run("Success Create", func(t *testing.T) {

		requestBody := []byte(`{"name": "Create Category Controller", "order": 10}`)

		request, e := http.NewRequest(
			"POST",
			"http://127.0.0.1:3000/api/categories",
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

func Test_GetAllCategory(t *testing.T) {
	Init()
	routes.Api(fiberApp)

	t.Run("Success Get All", func(t *testing.T) {

		request, e := http.NewRequest(
			"GET",
			"http://127.0.0.1:3000/api/categories",
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

func Test_GetCategoryById(t *testing.T) {
	Init()
	routes.Api(fiberApp)

	t.Run("Success Get by ID", func(t *testing.T) {
		categoryID := 1

		request, e := http.NewRequest(
			"GET",
			fmt.Sprintf("http://127.0.0.1:3000/api/categories/by-id/%d", categoryID),
			nil,
		)
		assert.Equal(t, nil, e)

		response, err := fiberApp.Test(request, -1)

		assert.Equal(t, nil, err)

		assert.Equal(t, fiber.StatusOK, response.StatusCode)

		b, err := io.ReadAll(response.Body)

		assert.Equal(t, nil, err)

		assert.JSONEq(t, `{"data":{"id":1,"created_at":"2024-04-25T13:04:23.626Z","updated_at":"2024-04-25T13:04:23.626Z","deleted_at":null,"name":"Category A","order":1},"message":"Success"}`, string(b))
	})
}

func Test_UpdateCategoryById(t *testing.T) {
	Init()
	routes.Api(fiberApp)

	t.Run("Success Update by ID", func(t *testing.T) {
		categoryID := 4

		requestBody := []byte(`{"name": "Updated Category New", "order": 20}`)

		request, e := http.NewRequest(
			"PUT",
			fmt.Sprintf("http://127.0.0.1:3000/api/categories/by-id/%d", categoryID),
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
		categoryID := 4

		request, e := http.NewRequest(
			"DELETE",
			fmt.Sprintf("http://127.0.0.1:3000/api/categories/by-id/%d", categoryID),
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
