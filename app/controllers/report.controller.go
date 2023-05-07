package controllers

import (
	"errors"
	"strconv"

	"github.com/devmeireles/go-vue-image-evaluate/app/models"
	"github.com/devmeireles/go-vue-image-evaluate/app/services"
	"github.com/devmeireles/go-vue-image-evaluate/app/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateReport godoc
// @Summary Endpoint to save avatar reports
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

// GetReport godoc
// @Summary Get details for a given report id
// @Description Get details of report corresponding to the input id
// @Tags report
// @Accept  json
// @Produce  json
// @Param id path int true "Report ID"
// @Success 200 {object} models.Report
// @Router /report/{id} [get]
func GetReport(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(res)
	}

	report, err := services.GetReportByID(id)

	if err != nil {
		res := utils.ResError(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(res)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(report)
	return c.Status(fiber.StatusOK).JSON(res)
}

// ListReports godoc
// @Summary Get a list of reports id
// @Description List all reports
// @Tags report
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Report
// @Router /report [get]
func ListReports(c *fiber.Ctx) error {
	reports, err := services.GetReports()

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(reports)
	return c.Status(fiber.StatusOK).JSON(res)
}
