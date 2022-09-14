package handlers

import (
	"gift-app/config"
	"gift-app/entities"
	"github.com/gofiber/fiber/v2"
)

func GetGift(c *fiber.Ctx) error {
	id := c.Params("id")
	var gift entities.Gift

	result := config.Database.Where("gift_id = ?", id).Find(&gift)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&gift)
}

func UseGift(c *fiber.Ctx) error {
	id := c.Params("id")

	gift := new(entities.Gift)
	config.Database.Where("gift_id = ?", id).First(&gift)

	gift.Amount = gift.Amount - 1

	if gift.Amount < 0 {
		return c.SendStatus(400)
	}

	config.Database.Save(&gift)

	return GetGift(c)
}
