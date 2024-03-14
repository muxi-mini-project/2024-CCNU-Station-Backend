package model

import "gorm.io/gorm"

type Building struct {
	gorm.Model
	BuildingName string `json:"building"`
}
