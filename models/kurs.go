package models

import (
	"github.com/google/uuid"
)

type Kurs struct {
	Id uuid.UUID `gorm:"type:uniqueidentifier;primary_key" json:"id"`
	Name string `gorm:"type:varchar(300)" json:"name"`
}