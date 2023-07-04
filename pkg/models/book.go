package models

import "gorm.io/gorm"

type Books struct {
	gorm.Model

	// ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
