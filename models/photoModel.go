package models

import "go-project/app"

type Photo struct {
	app.Base
	app.Photo

	UserID string
}
