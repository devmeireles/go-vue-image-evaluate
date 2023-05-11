package dto

type UpdateReportDTO struct {
	ExternalID string `json:"external_id" gorm:"text; not null;" validate:"required"`
	ImageUrl   string `json:"image_url" gorm:"text; not null;" validate:"required"`
	Status     int    `json:"status" gorm:"int; not null; default: 1"`
	Priority   int    `json:"priority" gorm:"int; not null; default: 1"`
}
