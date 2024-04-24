package routes

import (
	ans "sekolahbeta/final-project/question-random-generator/src/app/answer/controllers"
	cat "sekolahbeta/final-project/question-random-generator/src/app/category/controllers"
	que "sekolahbeta/final-project/question-random-generator/src/app/question/controllers"
	mod "sekolahbeta/final-project/question-random-generator/src/app/module/controllers"

	"github.com/gofiber/fiber/v2"
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

		answersGroup := route.Group("/answers")
		{
			answersGroup.Get("/", ans.GetAnswersList)
			answersGroup.Post("/", ans.InsertAnswerData)
			answersGroup.Get("/:id", ans.GetAnswerByID)
			answersGroup.Delete("/:id", ans.DeleteByID)
			answersGroup.Put("/:id", ans.UpdateAnswerByID)
		}

		modulesGroup := route.Group("/modules")
		{
			modulesGroup.Post("/", mod.InsertModuleData)
			modulesGroup.Get("/", mod.GetModulesList)
			modulesGroup.Get("/:id", mod.GetModuleByID)
			modulesGroup.Delete("/:id", mod.DeleteByID)
			modulesGroup.Put("/:id", mod.UpdateModuleByID)
			modulesGroup.Get("/exam/questions/:identifier", mod.GetQuestions)
		}
	}
}
