package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID uint `gorm:"primary_key" gen:"increment" not null json:"id"`

	CreatedAt time.Time  `gorm:"type:TIMESTAMP(0) with TIME ZONE;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:TIMESTAMP(0) with TIME ZONE;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:TIMESTAMP(0) with TIME ZONE;default:NULL" json:"deleted_at"`

	CreatedBy uint `gorm:"type:integer;default:0" json:"created_by"`
	UpdatedBy uint `gorm:"type:integer;default:0" json:"updated_by"`
	DeletedBy uint `gorm:"type:integer;default:0" json:"deleted_by"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("user_id")
	var userID uint
	if value != nil {
		userID = value.(uint)
	}
	m.CreatedBy = userID
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("user_id")
	var userID uint
	if value != nil {
		userID = value.(uint)
	}
	m.UpdatedBy = userID
	m.UpdatedAt = time.Now()
	return
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("user_id")
	var userID uint
	if value != nil {
		userID = value.(uint)
	}
	m.DeletedBy = userID
	now := time.Now()
	m.DeletedAt = &now
	return
}
