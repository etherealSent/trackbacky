package app

import (
	"awesomeProject/web-service-gin/internal/services"
    "awesomeProject/web-service-gin/internal/structs"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"regexp"
)

type DeezerResp struct {
	Data []Song `json:"data"`
}

type Song struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Link    string `json:"link"`
	Preview string `json:"preview"`
	Artist  struct {
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"artist"`
	Album struct {
		Cover string `json:"cover_xl"`
	} `json:"album"`
}

func GetTrack(c *gin.Context) {
	var requestData map[string]string
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	link, exists := requestData["url"]
	if !exists || link == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
		return
	}
	name, author := services.GetData(link)
	if name == "" || author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to extract song details from URL"})
		return
	}
	title := c.DefaultQuery("title", name)
	artist := c.DefaultQuery("artist", author)

	if title == "" || artist == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and artist are required"})
		return
	}

	encodedTitle := url.QueryEscape(title)
	encodedArtist := url.QueryEscape(artist)

	apiUrl := fmt.Sprintf("https://api.deezer.com/search?q=%s%%20%s", encodedTitle, encodedArtist)

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request: " + err.Error()})
		return
	}
	fmt.Println(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request: " + err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": fmt.Sprintf("Dezzer API returned status code: %d, body: %s", resp.StatusCode, resp)})
		return
	}
	var deezerResp DeezerResp
	err = json.NewDecoder(resp.Body).Decode(&deezerResp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response: " + err.Error()})
	}

	if len(deezerResp.Data) > 0 {
		firstTrack := deezerResp.Data[0]
		track := trackinfo.TrackInfo{
			Data: struct {
				Id              int32  `json:"id"`
				Title           string `json:"title"`
				Artist          string `json:"artist"`
				CoverUrl        string `json:"cover_url"`
				PreviewAudioUrl string `json:"preview"`
				Links           struct {
					DeezerUrl      string `json:"deezer_url"`
					SpotifyUrl     string `json:"spotify_url"`
					AppleMusicUrl  string `json:"apple_music_url"`
					YandexMusicUrl string `json:"yandex_music_url"`
					VkMusicUrl     string `json:"vk_music_url"`
				} `json:"links"`
			}{
				Id:              int32(firstTrack.ID),
				Title:           firstTrack.Title,
				Artist:          firstTrack.Artist.Name,
				CoverUrl:        firstTrack.Album.Cover,
				PreviewAudioUrl: firstTrack.Preview,
				Links: struct {
					DeezerUrl      string `json:"deezer_url"`
					SpotifyUrl     string `json:"spotify_url"`
					AppleMusicUrl  string `json:"apple_music_url"`
					YandexMusicUrl string `json:"yandex_music_url"`
					VkMusicUrl     string `json:"vk_music_url"`
				}{
					DeezerUrl:      firstTrack.Link,
					SpotifyUrl:     "qwe",
					AppleMusicUrl:  "rty",
					YandexMusicUrl: "asd",
					VkMusicUrl:     "zxc",
				},
			},
		}

		c.JSON(http.StatusOK, track)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "No Deezer track found"})
	}

}

type GetSongRequest struct {
	URL string `json:"url" example:"https://music.apple.com/song/697195787"`
}

func isValidURL(str string) bool {
	re := regexp.MustCompile(`^(https?://)?([\da-z.-]+)\.([a-z.]{2,6})([/\w .-]*)*/?$`)
	return re.MatchString(str)
}

// GetSong godoc
// @Summary      Получение информации о треке
// @Description  Возвращает информацию о музыкальном треке по переданной ссылке.
// @Tags         songs
// @Accept       json
// @Produce      json
// @Param        request body GetSongRequest true  "JSON URL трека"
// @Success      200   {object} trackinfo.TrackInfo "Информация о треке"
// @Failure      400   {object} trackinfo.ErrorResponse  "Некорректный запрос"
// @Failure      404   {object} trackinfo.ErrorResponse  "Трек не найден"
// @Router       /getsong [post]
func GetSong(c *gin.Context) {
	var requestData GetSongRequest

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

    if requestData.URL == "" {
        c.JSON(http.StatusBadRequest, trackinfo.ErrorResponse{Error: "URL is required"})
        return
    }

    if !isValidURL(requestData.URL) {
        c.JSON(http.StatusBadRequest, trackinfo.ErrorResponse{Error: "Invalid URL format"})
        return
    }

	fmt.Println("Полученный URL:", requestData.URL)

	c.JSON(200, gin.H{
		"data": map[string]any{
			"id":        3135556,
			"title":     "Harder, Better, Faster, Stronger",
			"artist":    "Daft Punk",
			"cover_url": "https://cdn-images.dzcdn.net/images/cover/5718f7c81c27e0b2417e2a4c45224f8a/1000x1000-000000-80-0-0.jpg",
			"preview":   "https://cdnt-preview.dzcdn.net/api/1/1/c/4/d/0/c4d7dbe3524ba59d2ad06d8cccd2484f.mp3?hdnea=exp=1738953605~acl=/api/1/1/c/4/d/0/c4d7dbe3524ba59d2ad06d8cccd2484f.mp3*~data=user_id=0,application_id=42~hmac=23b049d011241b86704326c18e9457f2f5218fe89b480890ff46345fb4dde757",
			"links": map[string]string{
				"deezer_url":       "https://www.deezer.com/track/3135556",
				"spotify_url":      "https://open.spotify.com/track/5W3cjX2J3tjhG8zb6u0qHn",
				"apple_music_url":  "https://music.apple.com/song/697195787",
				"yandex_music_url": "https://music.yandex.ru/track/270953",
				"vk_music_url":     "https://vk.com/audio-2001032355_114032355",
			},
		},
	})

}

func Ico(c *gin.Context) { c.JSON(204, gin.H{"favicon": nil}) }

//func DeezerData() {
//	rt := gin.Default()
//	url := ginSwagger.URL("http://0.0.0.0:8088/swagger/doc.json")
//	rt.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
//	rt.GET("/deezer", getDeezerTrack)
//	rt.GET("/favicon.ico", Ico)
//	rt.POST("/getsong", GetSong)
//	rt.Run("0.0.0.0:8088") // listen and serve on 0.0.0.0:8080
//}
