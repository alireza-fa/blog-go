package models

import (
	"database/sql"
	"time"
)

type BaseModel struct {
	Id int `gorm:"primarykey"`

	CreatedAt time.Time    `gorm:"type:TIMESTAMP with time zone;not null"`
	UpdatedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`
	DeletedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`

	CreatedBy int            `gorm:"not null"`
	UpdatedBy *sql.NullInt64 `gorm:"null"`
	DeletedBy *sql.NullInt64 `gorm:"null"`
}
