package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title     			string        `json:"title"`
	TodoStatus      string        `json:"todoStatus"`
	Description     string        `json:"description"`
	CreatedBy       string        `json:"createdby"`
}