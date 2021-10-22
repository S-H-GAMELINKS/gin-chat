package model

type Room struct {
	ID   uint64 `gorm:"primaryKey"`
	Name string `gorm:"type:string;not null"`
}
