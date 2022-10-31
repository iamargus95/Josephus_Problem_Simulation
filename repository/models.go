package models

import "gorm.io/gorm"

type Josephus struct {
	gorm.Model
	CircleOfDeath int `gorm:"unique"`
	LastAlive     int `gorm:"unique"`
}

func (t *Josephus) TableName() string {
	return "josephus"
}
