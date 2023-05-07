package dto

type CreateReportDTO struct {
	ReportID string `json:"report_id" gorm:"text; not null;" validate:"required"`
	Url      string `json:"url" gorm:"text; not null;" validate:"required"`
}
