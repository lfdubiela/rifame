package domain

import (
	"github.com/satori/go.uuid"
	"time"
)

type Model struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedBy uuid.UUID  `json:"created_by"`
	UpdatedBy uuid.UUID  `json:"modified_by"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}
