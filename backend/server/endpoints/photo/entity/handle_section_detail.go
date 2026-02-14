package photoEntity

import (
	"github.com/gofiber/fiber/v2"

	"share/types/model"

	"server/loaders/mysql"
	"server/procedures/photo"
	"server/types/common"
	"server/types/response"

	"share/types/payload"
)

func SectionDetailHandler(c *fiber.Ctx) error {
	// * Parse session
	s := c.Locals("session").(*common.Session)
	s.SetTag("photo/entity/view-section-detail")

	// * Parse query
	query := new(payload.SectionQuery)
	if err := c.QueryParser(query); err != nil {
		return response.Error(false, "Unable to parse query", err)
	}
	s.SetDetail("section-id", query.SectionId)

	// * Query section detail
	var photoSection *payload.ExtendedPhotoSection

	if result := mysql.DB.Model(new(model.PhotoSection)).
		Preload("PhotoAlbum").
		Select("photo_sections.*, (SELECT COUNT(*) FROM photo_items WHERE photo_section_id = photo_sections.id) AS photo_count").
		First(&photoSection, "id = ?", query.SectionId); result.Error != nil {
		return response.Error(false, "Unable to query album section", result.Error)
	}

	// * Check album permission
	if photoSection.PhotoAlbum.AccessToken != nil {
		if err := photo.CheckAlbumAccess(s, *photoSection.PhotoAlbum.Id, *photoSection.PhotoAlbum.AccessToken, query.Token); err != nil {
			_ = s.Commit(err)
			return err
		}
	}

	// * Generate photo access token
	// TODO: Implement JWT access token

	_ = s.Commit(nil)
	return c.JSON(response.Success(map[string]any{
		"section": &payload.AlbumSection{
			Id:           *photoSection.Id,
			Title:        *photoSection.Title,
			Subtitle:     *photoSection.Subtitle,
			Date:         *photoSection.Date,
			PhotoCount:   *photoSection.PhotoCount,
			ThumbnailUrl: "",
			UpdatedAt:    *photoSection.UpdatedAt,
		},
		"album": &payload.AlbumDetail{
			Id:       *photoSection.PhotoAlbum.Id,
			Name:     *photoSection.PhotoAlbum.Name,
			Slug:     *photoSection.PhotoAlbum.Slug,
			Sections: nil,
		},
	}))
}
