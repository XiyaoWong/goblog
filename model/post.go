// Package model provides ...
package model

import "github.com/jinzhu/gorm"

// Post ...
type Post struct {
	gorm.Model
	Title  string `gorm:"not null"`
	Body   string `gorm:"type:text"`
	IsShow int    `gorm:"default:1"`
}
