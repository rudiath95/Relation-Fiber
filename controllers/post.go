package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rudiath95/RelationFiber/ini"
	"github.com/rudiath95/RelationFiber/models"
)

func PostGetAll(c *fiber.Ctx) error {
	var posts []models.Post

	ini.DB.Preload("User").Find(&posts)

	return c.JSON(fiber.Map{
		"posts": posts,
	})
}

func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)

	//PARSE BODY REQUEST TO OBJECT STRUCT

	if err := c.BodyParser(post); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "can't handle request",
		})
	}

	//MANUAL VALIDATION
	if post.Title == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "title is required",
		})
	}
	if post.Body == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "body is required",
		})
	}

	if post.UserID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "user_id is required",
		})
	}

	ini.DB.Preload("User").Create(&post)

	return c.JSON(fiber.Map{
		"message": "Create Data Successfully",
		"post":    post,
	})
}
