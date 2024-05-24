package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	ID uuid.UUID `gorm:"primary_key;type:uuid" json:"id"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewRandom()
	base.ID = uuid

	return err
}
