package db

import (
	"fmt"
	"log"
	"video-microservice/pkg/config"
	"video-microservice/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initdb(cfg *config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.Db_host, cfg.Db_username, cfg.Db_password, cfg.Db_name, cfg.Db_port)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if dbErr != nil {
		log.Fatalln(dbErr)
	}
	db.AutoMigrate(&domain.Video{})

	return db, dbErr
}