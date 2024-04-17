package main

import (
	"sekolahbeta/final-project/question-random-generator/src/config"
	"sekolahbeta/final-project/question-random-generator/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func InitEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Warn("Cannot load env file, using system env")
	}
}

func main() {
	InitEnv()
	config.OpenDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5500",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	routes.Api(app)

	err := app.Listen(":3000")
	if err != nil {
		logrus.Fatal(
			"Error on running fiber, ",
			err.Error())
	}

}
