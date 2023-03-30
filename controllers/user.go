package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rudiath95/RelationFiber/ini"
	"github.com/rudiath95/RelationFiber/models"
)

func UserGetAll(c *fiber.Ctx) error {
	var users []models.User

	ini.DB.Preload("Locker").Find(&users)

	return c.JSON(fiber.Map{
		"users": users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	//PARSE BODY REQUEST TO OBJECT STRUCT

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "can't handle request",
		})
	}

	//MANUAL VALIDATION
	if user.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "name is required",
		})
	}

	ini.DB.Preload("Locker").Create(&user)

	return c.JSON(fiber.Map{
		"message": "Create Data Successfully",
		"user":    user,
	})
}
