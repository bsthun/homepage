package photoEntity

import (
	"github.com/gofiber/fiber/v2"

	"share/types/model"
	"share/utils/value"

	"server/loaders/mysql"
	"server/types/common"
	"server/types/response"

	"share/types/payload"
)

func AlbumListHandler(c *fiber.Ctx) error {
	// * Parse session
	s := c.Locals("session").(*common.Session)
	s.SetTag("photo/entity/view-album-list")

	// * Query photo detail
	var photoAlbums []*model.PhotoAlbum
	if result := mysql.DB.Find(&photoAlbums); result.Error != nil {
		return response.Error(false, "Unable to query photo albums", result.Error)
	}

	albums, _ := value.Iterate(photoAlbums, func(photoItem *model.PhotoAlbum) (*payload.AlbumList, *response.ErrorInstance) {
		return &payload.AlbumList{
			Id:            *photoItem.Id,
			Name:          *photoItem.Name,
			Slug:          *photoItem.Slug,
			CoverPhotoUrl: *photoItem.CoverPhotoUrl,
			SectionCount:  *photoItem.SectionCount,
			PhotoCount:    *photoItem.PhotoCount,
			UpdatedAt:     *photoItem.UpdatedAt,
		}, nil
	})

	_ = s.Commit(nil)
	return c.JSON(response.Success(map[string]any{
		"albums": albums,
	}))
}
