package photo

import (
	"server/loaders/mysql"
	"server/types/common"
	"server/types/response"

	"share/types/model"
)

func CheckAlbumAccess(session *common.Session, albumId uint64, albumToken string, queryToken string) *response.ErrorInstance {
	var count int
	if result := mysql.DB.
		Model(new(model.PhotoSessionAlbum)).
		Select("COUNT(*)").
		First(&count, "track_session_id = ? AND photo_album_id = ?", session.Id, albumId); result.Error != nil {
		return response.Error(true, "Unable to check album permission", result.Error)
	}

	if count == 0 {
		if albumToken != queryToken {
			return response.Error(false, "No album access permission")
		}

		sessionAlbum := &model.PhotoSessionAlbum{
			TrackSession:   nil,
			TrackSessionId: &session.Id,
			PhotoAlbum:     nil,
			PhotoAlbumId:   &albumId,
		}
		if result := mysql.DB.Create(sessionAlbum); result.Error != nil {
			return response.Error(true, "Unable to grant album permission to session", result.Error)
		}
	}

	return nil
}
