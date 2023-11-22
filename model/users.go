package model

import "github.com/google/uuid"

type Users struct {
	Id        uuid.UUID `gorm:"type:uuid;primary_key"`
	FullName  string    `gorm:"type:varchar(255)"`
	Email     string    `gorm:"type:varchar(255)"`
	Password  string    `gorm:"type:varchar(255)"`
	Phone     string    `gorm:"type:string"`
	Address   Address   `gorm:"embedded"`
	CreatedAt int64     `gorm:"autoCreateTime"`
	UpdatedAt int64     `gorm:"autoCreateTime"`
}
type Address struct {
	Region string `gorm:"type:varchar(255)"`
	City   string `gorm:"type:varchar(255)"`
}
