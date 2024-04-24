package controllers

import (
	"sekolahbeta/final-project/question-random-generator/src/app/answer/utils"
	"sekolahbeta/final-project/question-random-generator/src/app/answer/validation"
	"sekolahbeta/final-project/question-random-generator/src/app/models"

	"strconv"

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

	answer, errCreateAns := utils.InsertAnswerData(models.Answer{
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

func GetAnswerByID(c *fiber.Ctx) error {
	answerId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}

	answerData, err := utils.GetAnswersByID(uint(answerId))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
		logrus.Error("Error on get answer data: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    answerData,
			"message": "Success",
		},
	)
}

func UpdateAnswerByID(c *fiber.Ctx) error {
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

	answerId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}
	req.ID = answerId

	answerData, errUpdateData := utils.UpdateAnswersByID(models.Answer{
		Option:     req.Option,
		Answer:     req.Answer,
		Score:      req.Score,
		QuestionId: req.QuestionId,
	}, uint(req.ID))

	if errUpdateData != nil {
		logrus.Printf("Terjadi error : %s\n", errUpdateData.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"data":    answerData,
		"message": "Success Update Data",
	})
}

func DeleteByID(c *fiber.Ctx) error {
	answerId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}

	err = utils.DeleteByID(uint(answerId))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
		logrus.Error("Error on get answer data: ", err.Error())
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
