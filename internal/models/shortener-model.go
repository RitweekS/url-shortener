package models

import "gorm.io/gorm"

type Shortener struct {
	gorm.Model        // Includes fields like ID, CreatedAt, UpdatedAt, DeletedAt
	ShortStr   string `gorm:"unique;not null"`
	Url        string `gorm:"not null"`
}
