package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rudiath95/RelationFiber/ini"
	"github.com/rudiath95/RelationFiber/models"
)

func TagGetAll(c *fiber.Ctx) error {
	var tags []models.TagResponseWithPost

	ini.DB.Preload("Posts").Find(&tags)

	return c.JSON(fiber.Map{
		"tags": tags,
	})
}

func CreateTag(c *fiber.Ctx) error {
	tag := new(models.Tag)

	//PARSE BODY REQUEST TO OBJECT STRUCT

	if err := c.BodyParser(tag); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "can't handle request",
		})
	}

	//MANUAL VALIDATION
	if tag.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "name is required",
		})
	}

	ini.DB.Create(&tag)

	return c.JSON(fiber.Map{
		"message": "Create Data Successfully",
		"tag":     tag,
	})
}
