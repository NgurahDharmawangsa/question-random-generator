package routes

import (
	"github.com/gofiber/fiber/v2"
	cat "sekolahbeta/final-project/question-random-generator/src/app/category/controllers"
	que "sekolahbeta/final-project/question-random-generator/src/app/question/controllers"
)

func Api(app *fiber.App) {
	route := app.Group("api")
	{
		categoriesGroup := route.Group("/categories")
		{
			categoriesGroup.Get("/", cat.GetCategoriesList)
			categoriesGroup.Post("/", cat.InsertCategoryData)
			categoriesGroup.Get("/by-id/:id", cat.GetCategoryByID)
			categoriesGroup.Delete("/by-id/:id", cat.DeleteByID)
			categoriesGroup.Put("/by-id/:id", cat.UpdateCategoryByID)
		}

		questionsGroup := route.Group("/questions")
		{
			questionsGroup.Get("/", que.GetQuestionsList)
			questionsGroup.Post("/", que.InsertQuestionData)
			questionsGroup.Get("/by-id/:id", que.GetQuestionByID)
			questionsGroup.Delete("/by-id/:id", que.DeleteByID)
			questionsGroup.Put("/by-id/:id", que.UpdateQuestionByID)
		}
	}
}
