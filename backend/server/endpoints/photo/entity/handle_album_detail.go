package photoEntity

import (
	"github.com/gofiber/fiber/v2"

	model2 "share/types/model"
	value2 "share/utils/value"

	"server/loaders/mysql"
	"server/procedures/photo"
	"server/types/common"
	"server/types/response"

	"share/types/payload"
)

func AlbumDetailHandler(c *fiber.Ctx) error {
	// * Parse session
	s := c.Locals("session").(*common.Session)
	s.SetTag("photo/entity/view-album-detail")

	// * Parse query
	query := new(payload.AlbumQuery)
	if err := c.QueryParser(query); err != nil {
		return response.Error(false, "Unable to parse query", err)
	}
	s.SetDetail("album-slug", query.AlbumSlug)

	// * Query album detail
	var photoAlbum *model2.PhotoAlbum
	if result := mysql.DB.First(&photoAlbum, "slug = ?", query.AlbumSlug); result.Error != nil {
		return response.Error(false, "Unable to query album detail", result.Error)
	}

	// * Check album permission
	if photoAlbum.AccessToken != nil {
		if err := photo.CheckAlbumAccess(s, *photoAlbum.Id, *photoAlbum.AccessToken, query.Token); err != nil {
			_ = s.Commit(err)
			return err
		}
	}

	// * Query album sections
	var albumSections []*payload.ExtendedPhotoSection

	if result := mysql.DB.Model(new(model2.PhotoSection)).
		Select("photo_sections.*, (SELECT COUNT(*) FROM photo_items WHERE photo_section_id = photo_sections.id) AS photo_count, (SElECT GROUP_CONCAT(thumbnail_1) FROM (SELECT CONCAT(photo_sections.path, photo_items.thumbnail_path) AS thumbnail_1 FROM photo_items WHERE photo_items.photo_section_id = photo_sections.id ORDER BY RAND() LIMIT 10) T2) as thumbnail_url").
		Find(&albumSections, "photo_album_id = ?", photoAlbum.Id); result.Error != nil {
		return response.Error(true, "Unable to query list of album sections")
	}

	sections, _ := value2.Iterate(albumSections, func(albumSection *payload.ExtendedPhotoSection) (*payload.AlbumSection, *response.ErrorInstance) {
		return &payload.AlbumSection{
			Id:           *albumSection.Id,
			Title:        *albumSection.Title,
			Subtitle:     *albumSection.Subtitle,
			Date:         *albumSection.Date,
			PhotoCount:   *albumSection.PhotoCount,
			ThumbnailUrl: value2.Val(albumSection.ThumbnailUrl),
			UpdatedAt:    *albumSection.UpdatedAt,
		}, nil
	})

	_ = s.Commit(nil)
	return c.JSON(response.Success(map[string]any{
		"album": &payload.AlbumDetail{
			Id:       *photoAlbum.Id,
			Name:     *photoAlbum.Name,
			Slug:     *photoAlbum.Slug,
			Sections: sections,
		},
	}))
}
