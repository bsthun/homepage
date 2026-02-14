package hub

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"server/loaders/config"

	"share/types/extern"
)

func spotifyToken() error {
	// * Construct form data values
	values := url.Values{}
	values.Set("grant_type", "refresh_token")
	values.Set("refresh_token", config.C.SpotifyRefreshToken)

	// * Construct request
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+config.C.SpotifyAuthorization)

	// * Construct client
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// * Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// * Parse response body
	var data map[string]any
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	// * Set access token
	Hub.SpotifyAccessToken = data["access_token"].(string)
	Hub.SpotifyTokenExpired = time.Now().Add(time.Duration(data["expires_in"].(float64)) * time.Second)

	return nil
}

func spotifyPlayback() (*extern.SpotifyPlaybackState, error) {
	// * Construct request
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+Hub.SpotifyAccessToken)

	// * Construct client
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// * Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// * Check if body is isset
	if len(body) == 0 {
		return &extern.SpotifyPlaybackState{
			IsPlaying: false,
		}, nil
	}

	// * Parse response body
	var data *extern.SpotifyPlaybackState
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}
