package model

type Message struct {
	ID      uint64 `gorm:"primaryKey"`
	Content string `gorm:"type:text;not null"`
	RoomID  uint64 `gotm:"type:integer"`
}
