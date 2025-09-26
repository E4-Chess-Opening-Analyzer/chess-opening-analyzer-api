package models


type Game struct {
    ID     uint       `gorm:"primaryKey"`
    Result int        `gorm:"not null"` // 1: white win, 0: draw, -1: black win
    Moves  StringArray `json:"moves" gorm:"type:jsonb"`
}