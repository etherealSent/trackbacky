package models

import "gorm.io/gorm"

type Track struct {
	gorm.Model
	Id              int32  `gorm:"primaryKey"`
	Title           string `gorm:"not null"`
	Artist          string `gorm:"not null"`
	CoverUrl        string
	PreviewAudioUrl string
	DeezerUrl       string
	SpotifyUrl      string
	AppleMusicUrl   string
	YandexMusicUrl  string
	VkMusicUrl      string
}
