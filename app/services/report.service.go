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
