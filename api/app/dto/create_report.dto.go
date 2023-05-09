package dto

type CreateReportDTO struct {
	ExternalID string `json:"external_id" gorm:"text; not null;" validate:"required"`
	ImageUrl   string `json:"image_url" gorm:"text; not null;" validate:"required"`
}
