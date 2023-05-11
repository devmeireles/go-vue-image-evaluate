package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Report struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	ExternalID string    `json:"external_id" gorm:"text; not null;" validate:"required"`
	ImageUrl   string    `json:"image_url" gorm:"text; not null;" validate:"required"`
	Status     int       `json:"status" gorm:"int; not null; default: 1"`
	Priority   int       `json:"priority" gorm:"int; not null; default: 1"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime:true"`
}
