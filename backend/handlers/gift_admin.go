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

func CreateGift(c *fiber.Ctx) error {
	gift := new(entities.Gift)

	if err := c.BodyParser(gift); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Create(&gift)
	return c.Status(201).JSON(gift)
}

func UpdateGift(c *fiber.Ctx) error {
	id := c.Params("id")

	existingGift := new(entities.Gift)
	config.Database.Where("gift_id = ?", id).First(&existingGift)

	data := new(entities.Gift)

	if err := c.BodyParser(data); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	existingGift.Name = data.Name
	existingGift.DisplayText = data.DisplayText
	existingGift.Description = data.Description
	existingGift.Amount = data.Amount

	config.Database.Save(&existingGift)

	return GetGift(c)
}

func RemoveGift(c *fiber.Ctx) error {
	id := c.Params("id")
	var gift entities.Gift

	result := config.Database.Where("gift_id = ?", id).Delete(&gift)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
