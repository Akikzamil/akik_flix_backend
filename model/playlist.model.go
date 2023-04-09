package model

import "gorm.io/gorm"

type Playlist struct {
	gorm.Model
	Name string `json:"name"`
	Videos Video `gorm:"many2many:video_playlist;"`
}