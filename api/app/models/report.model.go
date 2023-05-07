package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Report struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;"`

	ReportID  string    `json:"report_id" gorm:"text; not null;" validate:"required"`
	Url       string    `json:"url" gorm:"text; not null;" validate:"required"`
	Status    int       `json:"status" gorm:"int; not null; default: 1"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:true"`
}