package services

import (
	database "github.com/devmeireles/go-vue-image-evaluate/app/config"
	"github.com/devmeireles/go-vue-image-evaluate/app/models"
)

func SaveReport(report *models.Report) (*models.Report, error) {
	if err := database.DB.Db.Model(&models.Report{}).Create(&report).Error; err != nil {
		return &models.Report{}, err
	}

	return report, nil
}

func GetReportByID(id int) (*models.Report, error) {
	var err error
	report := models.Report{}

	err = database.DB.Db.Model(&models.Report{}).
		First(&report, id).Error

	if err != nil {
		return &models.Report{}, err
	}

	return &report, nil
}

func GetReports() (*[]models.Report, error) {
	var err error
	reports := []models.Report{}

	err = database.DB.Db.Model(&models.Report{}).
		Find(&reports).Error

	if err != nil {
		return &[]models.Report{}, err
	}

	return &reports, nil
}
