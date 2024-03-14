package model

import (
	"gorm.io/gorm"
)

type GroupBasic struct {
	gorm.Model
	Name    string `json:"name"`
	OwnerID string `json:"owner_id"`
	Icon    string `json:"icon"`
	Desc    string `json:"desc"`
}
