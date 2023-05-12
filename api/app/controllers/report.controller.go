package controllers

import (
	"errors"
	"strconv"

	"github.com/devmeireles/go-vue-image-evaluate/app/dto"
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
// @Param evaluate body dto.CreateReportDTO true "Create a report"
// @Router /report [post]
func CreateReport(c *fiber.Ctx) error {
	data := new(dto.CreateReportDTO)

	c.BodyParser(data)

	errors := utils.ValidateStruct(*data)
	if errors != nil {
		res := utils.ResErrorValidation(errors)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	evaluateResponse, err := services.CreateEvaluateRequest(data.ImageUrl)

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	var classification = evaluateResponse.ClassificationOutcome
	var priority int

	switch classification {
	case "UnsafeContent_HighProbability", "RacyContent":
		priority = 3
		break
	case "SafeContent_ModerateProbability":
		priority = 2
	default:
		priority = 1
	}

	data.Priority = &priority

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
// @Param id path string true "Report ID"
// @Success 200 {object} models.Report
// @Router /report/{id} [get]
func GetReport(c *fiber.Ctx) error {
	id := c.Params("id")

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

	status, _ := strconv.Atoi(c.Query("status"))
	priority, _ := strconv.Atoi(c.Query("priority"))

	reports, err := services.GetReports(status, priority)

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(reports)
	return c.Status(fiber.StatusOK).JSON(res)
}

// UpdateReport godoc
// @Summary Update report identified by the given id
// @Description Update the report corresponding to the input id
// @Tags report
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "ID of the report to be updated"
// @Param report body dto.UpdateReportDTO true "Update report"
// @Success 200 {object} dto.UpdateReportDTO
// @Router /api/report/{id} [put]
func UpdateReport(c *fiber.Ctx) error {
	id := c.Params("id")
	payload := new(dto.UpdateReportDTO)
	c.BodyParser(payload)

	validateErrors := utils.ValidateStruct(*payload)
	if validateErrors != nil {
		res := utils.ResErrorValidation(validateErrors)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	_, err := services.GetReportByID(id)

	if err != nil {
		res := utils.ResError(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(res)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	report, err := services.UpdateReport(payload, id)
	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(report)
	return c.Status(fiber.StatusOK).JSON(res)
}
