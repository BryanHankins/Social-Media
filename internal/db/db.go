package db

import (
    "log"

    "github.com/yourusername/socialmedia/internal/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    dsn := "host=localhost user=postgres password=yourpassword dbname=socialmedia port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
}

func Migrate() {
    DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Like{}, &models.Follow{}, &models.Notification{})
}
