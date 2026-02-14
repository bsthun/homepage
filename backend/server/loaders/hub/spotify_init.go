package hub

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"

	"share/types/enum"
	"share/utils/text"
)

func SpotifyInit() {
	if err := spotifyToken(); err != nil {
		logrus.Warn("Unable to get Spotify token: ", err)
	}

	s := gocron.NewScheduler(text.BangkokTime)
	_, _ = s.Every(1 * time.Minute).Do(func() {
		if time.Now().After(Hub.SpotifyTokenExpired) {
			if err := spotifyToken(); err != nil {
				logrus.Warn("Unable to refresh Spotify token: ", err)
			}
		}
		state, err := spotifyPlayback()
		if err != nil {
			logrus.Warn("Unable to get Spotify playback state: ", err)
			return
		}

		// * Check if playback is active
		if !state.IsPlaying || state.Item == nil {
			return
		}

		// * Check if playback is a track
		if state.Item.Type != "track" {
			return
		}

		// * Check if track is registered
		spotifyRecord(enum.SpotifyEntityTrack, state.Item.Id)
		for _, artist := range state.Item.Artists {
			spotifyRecord(enum.SpotifyEntityArtist, artist.Id)
		}
		if state.Context != nil {
			spotifyRecord(enum.SpotifyEntityContext, state.Context.Uri)
		}
		spotifyRecord(enum.SpotifyEntityDevice, state.Device.Name)
		spotifyRecord(enum.SpotifyEntityRepeat, state.RepeatState)
	})
	s.StartAsync()
}
