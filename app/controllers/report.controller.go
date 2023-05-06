package controllers

import (
	"github.com/devmeireles/go-vue-image-evaluate/app/models"
	"github.com/devmeireles/go-vue-image-evaluate/app/services"
	"github.com/devmeireles/go-vue-image-evaluate/app/utils"
	"github.com/gofiber/fiber/v2"
)

// CreateReport godoc
// @Summary Endpoint to saves avatar reports
// @Description Saves a request for avatar evaluate
// @Tags report
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Report
// @Param evaluate body models.Report true "Create a report"
// @Router /report [post]
func CreateReport(c *fiber.Ctx) error {
	data := new(models.Report)

	c.BodyParser(data)

	errors := utils.ValidateStruct(*data)
	if errors != nil {
		res := utils.ResErrorValidation(errors)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	report, err := services.SaveReport(data)

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(report)
	return c.Status(fiber.StatusCreated).JSON(res)
}
