package model

type Room struct {
	ID       uint64 `gorm:"primary_key"`
	Name     string `gorm:"type:string;not null"`
	Messages *[]Message
}
