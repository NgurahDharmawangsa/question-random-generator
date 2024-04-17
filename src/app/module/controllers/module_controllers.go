package controllers

import (
	"fmt"
	"math/rand"
	"sekolahbeta/final-project/question-random-generator/src/app/module/model"
	mod "sekolahbeta/final-project/question-random-generator/src/app/module/utils"
	"sekolahbeta/final-project/question-random-generator/src/app/module/validation"
	que "sekolahbeta/final-project/question-random-generator/src/app/question/utils"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
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

	questionsData, err := que.GetQuestionsList()
	if err != nil {
		logrus.Error("Error on get Questions list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}

	var ids []uint
	for _, question := range questionsData {
		ids = append(ids, question.ID)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ids), func(i, j int) {
		ids[i], ids[j] = ids[j], ids[i]
	})
	fmt.Println(ids)

	// Konversi dari []uint ke pq.Int64Array
	var pqIds pq.Int64Array
	for _, id := range ids {
		pqIds = append(pqIds, int64(id))
	}

	module, errCreateMod := mod.InsertModuleData(model.Module{
		Identifier:  "MDL-" + time.Now().Format("20060102150405"),
		Name:        req.Name,
		QuestionIds: pqIds,
	})

	if errCreateMod != nil {
		logrus.Printf("Terjadi error : %s\n", errCreateMod.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",
			})
	}

	return c.Status(fiber.StatusCreated).JSON(map[string]any{
		"data":    module,
		"message": "Success Insert Data",
	})

}
