package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Name string `json:"name"`
	Path string `json:"path"`
	Playlists []Playlist `gorm:"many2many:video_playlist;"`
}