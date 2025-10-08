// models/models.go
package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey"`
	Description string `json:"description" gorm:"text"`
	Note        string `json:"note" gorm:"text"`
}
