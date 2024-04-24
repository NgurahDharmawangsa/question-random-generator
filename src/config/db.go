package config

import (
	"fmt"
	"log"
	"os"
	"sekolahbeta/final-project/question-random-generator/src/app/models"

	// ans "sekolahbeta/final-project/question-random-generator/src/app/answer/model"
	// cat "sekolahbeta/final-project/question-random-generator/src/app/category/model"
	// que "sekolahbeta/final-project/question-random-generator/src/app/question/model"
	// mod "sekolahbeta/final-project/question-random-generator/src/app/module/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct {
	DB *gorm.DB
}

var Mysql MysqlDB

func OpenDB() {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	mysqlConn, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Mysql = MysqlDB{
		DB: mysqlConn,
	}

	err = autoMigrate(mysqlConn)
	if err != nil {
		log.Fatal(err)
	}

}

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Category{}, &models.Question{}, &models.Answer{}, &models.Module{},
	)

	if err != nil {
		return err
	}

	return nil
}
