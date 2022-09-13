package handlers

import (
	"gift-app/config"
	"gift-app/entities"
	"github.com/gofiber/fiber/v2"
)

func GetGifts(c *fiber.Ctx) error {
	var gifts []entities.Gift

	config.Database.Find(&gifts)
	return c.Status(200).JSON(gifts)
}

func GetGift(c *fiber.Ctx) error {
	id := c.Params("id")
	var gift entities.Gift

	result := config.Database.Find(&gift, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&gift)
}

func AddGift(c *fiber.Ctx) error {
	gift := new(entities.Gift)

	if err := c.BodyParser(gift); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&gift)
	return c.Status(201).JSON(gift)
}

func UpdateGift(c *fiber.Ctx) error {
	gift := new(entities.Gift)
	id := c.Params("id")

	if err := c.BodyParser(gift); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("ID = ?", id).Updates(&gift)

	return GetGift(c)
	//return c.Status(200).JSON(gift)
}

func RemoveGift(c *fiber.Ctx) error {
	id := c.Params("id")
	var gift entities.Gift

	result := config.Database.Delete(&gift, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
