package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"hw8/models"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

// ConnectDb инициализирует соединение с базой данных
func ConnectDb() {
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("Ошибка при получении пути к исполняемому файлу:", err)
		return
	}

	// Получение абсолютного пути к каталогу с исполняемым файлом
	currentDir := filepath.Dir(executablePath)

	fmt.Println("Текущий рабочий каталог в database.go:", currentDir)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Astana",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to the database.\n", err)
		os.Exit(2)
	}

	log.Println("Connected to the database")

	// Автомиграция создаст таблицы на основе моделей
	db.AutoMigrate(&models.Fact{})

	DB = DbInstance{
		Db: db,
	}
}

// Close закрывает соединение с базой данных
func Close() {
	db, err := DB.Db.DB()
	if err != nil {
		log.Fatal("Failed to close the database connection.\n", err)
		os.Exit(2)
	}
	db.Close()
}
