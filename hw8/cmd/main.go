package main

import (
	"fmt"
	"hw8/cmd"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/hw8/music-library-api/database"
)

func main() {
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("Ошибка при получении пути к исполняемому файлу:", err)
		return
	}

	// Получение абсолютного пути к каталогу с исполняемым файлом
	currentDir := filepath.Dir(executablePath)

	fmt.Println("Текущий рабочий каталог:", currentDir)

	app := fiber.New()

	// Инициализация соединения с базой данных
	database.ConnectDb()

	// Настройка маршрутов
	cmd.SetupRoutes(app)

	// Запуск сервера
	app.Listen(":3000")
}
