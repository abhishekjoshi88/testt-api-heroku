package main

import (
	"log"

	"cleanarch/config"
	Delivery "cleanarch/delivery"
	Repository "cleanarch/repository"
	UseCase "cleanarch/usecase"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	e *echo.Echo
)

func init() {
	// Initialise echo context for routes
	e = echo.New()

	config.InitializeConfig()
}

func main() {

	// Load database config from config.yml/environment
	databaseConfig, err := config.GetDbConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Establish data base connection
	db, err := gorm.Open(mysql.Open(databaseConfig.DbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		e.Logger.Fatal(err)
	}

	repo := Repository.StickerConstructor(db)
	usecase := UseCase.StickerConstructor(repo)
	Delivery.StickerConstructor(e, usecase)

	port := viper.GetString("APPLICATION_PORT")

	log.Fatal(e.Start(":" + port))

}
