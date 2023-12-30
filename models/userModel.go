package models

import "go-project/app"

type User struct {
	app.Base
	app.User
	Photos []Photo `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"photos"`
}
