package hub

import (
	"time"

	"server/loaders/config"
)

var Hub *hubModel

func Init() {
	Hub = &hubModel{
		SpotifyAccessToken:  "",
		SpotifyTokenExpired: time.Now(),
	}

	if config.C.SpotifyAuthorization != "" {
		SpotifyInit()
	}
}
