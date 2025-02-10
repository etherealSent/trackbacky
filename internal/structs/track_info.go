package track_info

type TrackInfo struct {
	Data struct {
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
	} `json:"data"`
}
