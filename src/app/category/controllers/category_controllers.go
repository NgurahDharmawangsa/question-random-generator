package controllers

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"sekolahbeta/final-project/question-random-generator/src/app/category/model"
	"sekolahbeta/final-project/question-random-generator/src/app/category/utils"
	"sekolahbeta/final-project/question-random-generator/src/app/category/validation"
	"strconv"
)

func InsertCategoryData(c *fiber.Ctx) error {
	req := new(validation.AddCategoryRequest)

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

	category, errCreateCar := utils.InsertCategoryData(model.Category{
		Name:  req.Name,
		Order: req.Order,
	})

	if errCreateCar != nil {
		logrus.Printf("Terjadi error : %s\n", errCreateCar.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"data":    category,
		"message": "Success Insert Data",
	})
}

func GetCategoriesList(c *fiber.Ctx) error {
	categoriesData, err := utils.GetCategoriesList()
	if err != nil {
		logrus.Error("Error on get categories list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    categoriesData,
			"message": "Success",
		},
	)
}

func GetCategoryByID(c *fiber.Ctx) error {
	categoryId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}

	categoryData, err := utils.GetCategoryByID(uint(categoryId))
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
			"data":    categoryData,
			"message": "Success",
		},
	)
}

func UpdateCategoryByID(c *fiber.Ctx) error {
	type AddCategoryRequest struct {
		ID    int    `json:"id" form:"id"`
		Name  string `json:"name" valid:"required"`
		Order string `json:"order" valid:"required"`
	}

	req := new(AddCategoryRequest)

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

	categoryId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}
	req.ID = categoryId

	categoryData, errUpdateData := utils.UpdateCategoriesByID(model.Category{
		Name:  req.Name,
		Order: req.Order,
	}, uint(req.ID))

	if errUpdateData != nil {
		logrus.Printf("Terjadi error : %s\n", errUpdateData.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"data":    categoryData,
		"message": "Success Update Data",
	})
}

func DeleteByID(c *fiber.Ctx) error {
	categoryId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "ID not valid",
			},
		)
	}

	err = utils.DeleteByID(uint(categoryId))
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
