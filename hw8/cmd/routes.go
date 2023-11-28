package cmd

import (
	"fmt"
	"hw8/handlers"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("Ошибка при получении пути к исполняемому файлу:", err)
		return
	}

	// Получение абсолютного пути к каталогу с исполняемым файлом
	currentDir := filepath.Dir(executablePath)

	fmt.Println("Текущий рабочий каталог в routes.go:", currentDir)

	// Роуты для фактов (facts)
	app.Get("/facts", handlers.ListFacts)
	app.Get("/facts/:id", handlers.GetFact)
	app.Post("/facts", handlers.CreateFact)
	app.Put("/facts/:id", handlers.UpdateFact)
	app.Delete("/facts/:id", handlers.DeleteFact)

	// ... (ваш код для треков и плейлистов)
}
