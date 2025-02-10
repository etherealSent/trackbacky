package transport

import (
	"awesomeProject/web-service-gin/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSong(c *gin.Context) {

	var request struct {
		DeezerURL string `json:"DeezerURL"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный запрос"})
		return
	}

	track, err := repositories.GetTrackByDeezer(request.DeezerURL)
	if err == nil {
		c.JSON(http.StatusOK, track)
		return
	}

	newTrack := parseTrackFromDeezer(request.DeezerURL)
	err = repositories.SaveTrack(&newTrack)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка записи в бд"})
	}

	c.JSON(http.StatusOK, newTrack)
}
