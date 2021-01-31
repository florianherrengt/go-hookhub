package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// HookEvent from external service
type HookEvent struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primary_key;"`
	Payload string
}

// BeforeCreate adds the uuid
func (i *HookEvent) BeforeCreate(tx *gorm.DB) (err error) {
	i.ID = uuid.New()
	return
}
