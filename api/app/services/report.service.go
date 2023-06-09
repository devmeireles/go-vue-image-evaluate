package services

import (
	database "github.com/devmeireles/go-vue-image-evaluate/app/config"
	"github.com/devmeireles/go-vue-image-evaluate/app/dto"
	"github.com/devmeireles/go-vue-image-evaluate/app/models"
	uuid "github.com/satori/go.uuid"
)

func SaveReport(report *dto.CreateReportDTO) (*models.Report, error) {
	newReport := &models.Report{
		ID:         uuid.NewV4(),
		ImageUrl:   report.ImageUrl,
		ExternalID: report.ExternalID,
		Priority:   int(*report.Priority),
	}

	if err := database.DB.Db.Model(&models.Report{}).Create(newReport).Error; err != nil {
		return &models.Report{}, err
	}

	return newReport, nil
}

func GetReportByID(id string) (*models.Report, error) {
	var err error
	report := models.Report{}

	err = database.DB.Db.Model(&models.Report{}).
		First(&report, "id = ?", id).Error

	if err != nil {
		return &models.Report{}, err
	}

	return &report, nil
}

func GetReports(status int, priority int) (*[]models.Report, error) {
	var err error
	reports := []models.Report{}

	query := database.DB.Db.Model(&models.Report{}).
		Order("status ASC").Order("priority DESC")

	if status != 0 {
		query = query.Where("status = ?", status)
	}

	if priority != 0 {
		query = query.Where("priority = ?", priority)
	}

	err = query.Find(&reports).Error

	if err != nil {
		return &[]models.Report{}, err
	}

	return &reports, nil
}

func UpdateReport(report *dto.UpdateReportDTO, id string) (*models.Report, error) {
	updateReport := &models.Report{
		ImageUrl:   report.ImageUrl,
		ExternalID: report.ExternalID,
		Status:     report.Status,
		Priority:   report.Priority,
	}

	if err := database.DB.Db.Model(&models.Report{}).
		Where("id = ?", id).
		Updates(&updateReport).Error; err != nil {
		return &models.Report{}, err
	}

	return updateReport, nil
}
