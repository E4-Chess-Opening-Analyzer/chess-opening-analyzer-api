package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model `swaggerignore:"true"`
	Result	 int `gorm:"not null"` // 1: white win, 0: draw, -1: black win
	Moves	 []string `json:"moves" gorm:"type:text[]"`
}