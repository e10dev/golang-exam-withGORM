package model

type User struct {
	Seq         int    `gorm:"primaryKey;" json:"seq"`
	Id          string `gorm:"not null" json:"id"`
	Pw          string `gorm:"not null" json:"pw"`
	Name        string `gorm:"not null" json:"name"`
	Email       string `gorm:"not null" json:"email"`
	Hp          string `json:"hp"`
	Role        int    `gorm:"not null" json:"role"`
	State       int    `gorm:"not null" json:"state"`
	Description string `json:"description"`
}

type UserCompact struct {
	Id string `json:"id"`
	Pw string `json:"pw"`
}

var Seq int
