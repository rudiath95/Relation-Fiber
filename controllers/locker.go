package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rudiath95/RelationFiber/ini"
	"github.com/rudiath95/RelationFiber/models"
)

func LockerGetAll(c *fiber.Ctx) error {
	var lockers []models.Locker

	ini.DB.Preload("User").Find(&lockers)

	return c.JSON(fiber.Map{
		"lockers": lockers,
	})
}

func CreateLocker(c *fiber.Ctx) error {
	locker := new(models.Locker)

	//PARSE BODY REQUEST TO OBJECT STRUCT

	if err := c.BodyParser(locker); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "can't handle request",
		})
	}

	//MANUAL VALIDATION
	if locker.Code == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "code is required",
		})
	}

	if locker.UserID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "user_id is required",
		})
	}

	ini.DB.Preload("User").Create(&locker)

	return c.JSON(fiber.Map{
		"message": "Create Data Successfully",
		"locker":  locker,
	})
}
