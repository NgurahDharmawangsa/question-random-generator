package controllers

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	cat "sekolahbeta/final-project/question-random-generator/src/app/category/model"
	"sekolahbeta/final-project/question-random-generator/src/app/question/model"
	"sekolahbeta/final-project/question-random-generator/src/app/question/utils"
	"sekolahbeta/final-project/question-random-generator/src/app/question/validation"
	"strconv"
)

func InsertQuestionData(c *fiber.Ctx) error {
	req := new(validation.AddQuestionRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{
				"message": "Body not valid",
			})
	}

	isValid, err := govalidator.ValidateStruct(req)
	if !isValid && err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": err.Error(),
		})
	}

	question, errCreateCar := utils.InsertQuestionData(model.Question{
		Question:   req.Question,
		CategoryId: req.CategoryId,
	})

	if errCreateCar != nil {
		logrus.Printf("Terjadi error : %s\n", errCreateCar.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"data":    question,
		"message": "Success Insert Data",
	})
}

func GetQuestionsList(c *fiber.Ctx) error {
	questionsData, err := utils.GetQuestionsList()
	if err != nil {
		logrus.Error("Error on get Questions list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    questionsData,
			"message": "Success",
		},
	)
}

func GetQuestionByID(c *fiber.Ctx) error {
	questionId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}

	questionData, err := utils.GetQuestionsByID(uint(questionId))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
		logrus.Error("Error on get car data: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    questionData,
			"message": "Success",
		},
	)
}

func UpdateQuestionByID(c *fiber.Ctx) error {
	req := new(validation.AddQuestionRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{
				"message": "Body not valid",
			})
	}

	isValid, err := govalidator.ValidateStruct(req)
	if !isValid && err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]any{
			"message": err.Error(),
		})
	}

	questionId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}
	req.ID = questionId

	questionData, errUpdateData := utils.UpdateQuestionsByID(model.Question{
		Question:   req.Question,
		CategoryId: req.CategoryId,
		Category:   cat.Category{},
	}, uint(req.ID))

	if errUpdateData != nil {
		logrus.Printf("Terjadi error : %s\n", errUpdateData.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"data":    questionData,
		"message": "Success Update Data",
	})
}

func DeleteByID(c *fiber.Ctx) error {
	questionId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}

	err = utils.DeleteByID(uint(questionId))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
		logrus.Error("Error on get car data: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"message": "Success Delete Data",
		},
	)
}
