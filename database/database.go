package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Report struct {
	ReportID     int64     `json:"report_id" gorm:"type:int;primaryKey"`
	CreationTime time.Time `json:"creation_time" gorm:"type:date"`
	ReportInfo   string    `json:"report_info" gorm:"type:char"`
	ModelID      int64     `json:"model_id" gorm:"type:int"`
}

func CreateTable() {
	db := DBConnect()
	db.AutoMigrate(&Report{})
}

func DBConnect() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
