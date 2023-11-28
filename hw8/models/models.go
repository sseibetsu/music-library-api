package models

import (
	"time"

	"gorm.io/gorm"
)

type Track struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(255);not null"`
	Artist      string `json:"artist" gorm:"type:varchar(255);not null"`
	Album       string `json:"album" gorm:"type:varchar(255);not null"`
	ReleaseYear int    `json:"release_year" gorm:"not null"`
	// Другие поля, соответствующие вашим требованиям...
}

type Playlist struct {
	gorm.Model
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" gorm:"text"`
	Tracks      []Track   `json:"tracks" gorm:"many2many:playlist_tracks;"`
	CreatedAt   time.Time `json:"created_at"`
	// Другие поля, соответствующие вашим требованиям...
}
