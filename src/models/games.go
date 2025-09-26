package models


type Game struct {
	ID 		 uint `gorm:"primarykey"`
	Result	 int `gorm:"not null"` // 1: white win, 0: draw, -1: black win
	Moves	 []string `json:"moves" gorm:"type:text[]"`
}