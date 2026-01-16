package resource

import (
	"time"

	"gorm.io/plugin/optimisticlock"
)

type resource struct {
	ID           uint
	Version      optimisticlock.Version
	CreatedBy    uint
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	LastUpdated  time.Time `gorm:"autoUpdateTime"`
	TextField    *string   `gorm:"size:255"`
	NumberField  *int
	BooleanField *bool
}

func (resource) TableName() string {
	return "public.resources"
}
