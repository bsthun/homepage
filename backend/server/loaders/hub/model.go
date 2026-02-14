package hub

import "time"

type hubModel struct {
	SpotifyAccessToken  string
	SpotifyTokenExpired time.Time
}
