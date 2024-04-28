package controllers

import (
	"sekolahbeta/final-project/question-random-generator/src/app/models"
	mod "sekolahbeta/final-project/question-random-generator/src/app/module/utils"
	"sekolahbeta/final-project/question-random-generator/src/app/module/validation"

	// que "sekolahbeta/final-project/question-random-generator/src/app/question/utils"

	"strconv"

	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func InsertModuleData(c *fiber.Ctx) error {
	req := new(validation.AddModuleRequest)

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

	pst := models.Module{
		Identifier:  "MDL-" + time.Now().Format("20060102150405"),
		Name:        req.Name,
		QuestionIds: req.QuestionIds,
	}

	module, errCreateMod := mod.InsertModuleData(pst)

	if errCreateMod != nil {
		logrus.Printf("Terjadi error : %s\n", errCreateMod.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}

	res := validation.AddModuleRequest{
		Name:        module.Name,
		QuestionIds: module.QuestionIds,
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"data":    res,
		"message": "Success Insert Data",
	})

}

func GetModulesList(c *fiber.Ctx) error {
	modulesData, err := mod.GetModulesList()
	if err != nil {
		logrus.Error("Error on get Modules list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    modulesData,
			"message": "Success",
		},
	)
}

func GetModuleByID(c *fiber.Ctx) error {
	moduleId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}

	moduleData, err := mod.GetModulesByID(uint(moduleId))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
		logrus.Error("Error on get Module data: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    moduleData,
			"message": "Success",
		},
	)
}

func UpdateModuleByID(c *fiber.Ctx) error {
	req := new(validation.AddModuleRequest)

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

	moduleId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}
	req.ID = moduleId

	moduleData, errUpdateData := mod.UpdateModulesByID(models.Module{
		Identifier:  req.Identifier,
		Name:        req.Name,
		QuestionIds: req.QuestionIds,
	}, uint(req.ID))

	if errUpdateData != nil {
		logrus.Printf("Terjadi error : %s\n", errUpdateData.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"data":    moduleData,
		"message": "Success Update Data",
	})
}

func DeleteByID(c *fiber.Ctx) error {
	moduleId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}

	err = mod.DeleteByID(uint(moduleId))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",
				},
			)
		}
		logrus.Error("Error on get module data: ", err.Error())
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

func GetQuestions(c *fiber.Ctx) error {
	moduleIdentifier := c.Params("identifier")
	if moduleIdentifier == "" {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]interface{}{
				"message": "ID tidak valid",
			},
		)
	}

	moduleData, err := mod.GetQuestions(moduleIdentifier)
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "Identifier not found",
				},
			)
		}
		logrus.Error("Error on get Module data: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}

	type QuestionData struct {
		Identifier   string            `json:"identifier"`
		Name         string            `json:"name"`
		ExamQuestion []models.Question `json:"exam_question"`
	}

	res := QuestionData{
		Identifier:   moduleData.Identifier,
		Name:         moduleData.Name,
		ExamQuestion: moduleData.Question,
	}

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    res,
			"message": "Success",
		},
	)
}
