package handlers

import (
	"fmt"
	"hw8/database"
	"hw8/models"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("Ошибка при получении пути к исполняемому файлу:", err)
		return err
	}

	// Получение абсолютного пути к каталогу с исполняемым файлом
	currentDir := filepath.Dir(executablePath)

	fmt.Println("Текущий рабочий каталог в facts.go:", currentDir)

	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

// ListTracks возвращает список всех треков
func ListTracks(c *fiber.Ctx) error {
	tracks := []models.Track{}
	database.DB.Db.Find(&tracks)

	return c.Status(fiber.StatusOK).JSON(tracks)
}

// GetTrack возвращает информацию о конкретном треке
func GetTrack(c *fiber.Ctx) error {
	track := models.Track{}
	if err := database.DB.Db.First(&track, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Track not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(track)
}

// CreateTrack создает новый трек
func CreateTrack(c *fiber.Ctx) error {
	track := new(models.Track)
	if err := c.BodyParser(track); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&track)

	return c.Status(fiber.StatusOK).JSON(track)
}

// UpdateTrack обновляет информацию о треке
func UpdateTrack(c *fiber.Ctx) error {
	track := models.Track{}
	if err := database.DB.Db.First(&track, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Track not found",
		})
	}

	if err := c.BodyParser(&track); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Save(&track)

	return c.Status(fiber.StatusOK).JSON(track)
}

// DeleteTrack удаляет трек
func DeleteTrack(c *fiber.Ctx) error {
	track := models.Track{}
	if err := database.DB.Db.First(&track, c.Params("id")).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Track not found",
		})
	}

	database.DB.Db.Delete(&track)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Track deleted successfully",
	})
}
