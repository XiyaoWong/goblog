// Package model provides ...
package model

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Title string
	Body  string
}