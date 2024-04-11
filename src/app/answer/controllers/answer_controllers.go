package controllers

import (
	"sekolahbeta/final-project/question-random-generator/src/app/answer/model"
	"sekolahbeta/final-project/question-random-generator/src/app/answer/utils"
	"sekolahbeta/final-project/question-random-generator/src/app/answer/validation"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func InsertAnswerData(c *fiber.Ctx) error {
	req := new(validation.AddAnswerRequest)

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

	answer, errCreateAns := utils.InsertAnswerData(model.Answer{
		Option:     req.Option,
		Answer:     req.Answer,
		Score:      req.Score,
		QuestionId: req.QuestionId,
	})

	if errCreateAns != nil {
		logrus.Printf("Terjadi error : %s\n", errCreateAns.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"data":    answer,
		"message": "Success Insert Data",
	})
}

func GetAnswersList(c *fiber.Ctx) error {
	answersData, err := utils.GetAnswersList()
	if err != nil {
		logrus.Error("Error on get Answers list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    answersData,
			"message": "Success",
		},
	)
}
