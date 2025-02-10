package repositories

import (
	"awesomeProject/web-service-gin/internal/database"
	"awesomeProject/web-service-gin/internal/models"
)

func GetTrackByDeezer(url string) (*models.Track, error) {
	var track models.Track
	err := database.DB.Where("DeezerUrl = ?", url).First(&track).Error
	if err != nil {
		return nil, err
	}
	return &track, nil
}

func SaveTrack(track *models.Track) error {
	return database.DB.Save(track).Error
}
