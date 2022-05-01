package handler

import (
	"api-url-shortener/database"
	"api-url-shortener/model"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAllLinks(c *fiber.Ctx) error {
	var links []model.Link

	result := database.DB.Find(&links)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return c.JSON(links)
}

func RedirectByShortenedLink(c *fiber.Ctx) error {
	shortlink := c.Params("shortened_link")

	var link model.Link
	err := database.DB.Where("shortened_link=?", shortlink).Find(&link).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "link not found",
		})
	}
	if link.Id == 0 || link.OriginLink == "" {
		return c.Status(404).JSON(fiber.Map{
			"message": "link not found",
		})
	}
	return c.Redirect(link.OriginLink)
}

func CreateLink(c *fiber.Ctx) error {
	link := new(model.CreateLink)
	if err := c.BodyParser(link); err != nil {
		return err
	}

	newLink := model.Link{
		OriginLink:    link.OriginLink,
		ShortenedLink: link.ShortenedLink,
	}

	errCreateLink := database.DB.Create(&newLink).Error
	if errCreateLink != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed create link",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    newLink,
	})
}

func UpdateByShortenedLink(c *fiber.Ctx) error {
	linkRequest := new(model.UpdateLink)
	if err := c.BodyParser(linkRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var link model.Link

	shortlink := c.Params("shortened_link")

	//CHECK AVAILABLE LINK
	err := database.DB.Where("shortened_link=?", shortlink).Find(&link).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "link not found",
		})
	}

	//UPDATE USER DATA
	if linkRequest.OriginLink != "" {
		link.OriginLink = linkRequest.OriginLink
	}
	errUpdate := database.DB.Save(&link).Error
	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    link,
	})
}

func DeleteByShortenedLink(c *fiber.Ctx) error {
	shortlink := c.Params("shortened_link")
	var link model.Link

	//CHECK AVAILABLE LINK
	err := database.DB.Where("shortened_link=?", shortlink).Find(&link).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "link not found",
		})
	}

	errDelete := database.DB.Delete(&link).Error
	if errDelete != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "link deleted",
	})
}
