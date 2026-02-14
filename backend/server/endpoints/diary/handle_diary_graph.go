package diaryEndpoints

import (
	"fmt"
	"time"

	strapiModel "share/types/model_strapi"
	"share/utils/text"
	"share/utils/value"

	"github.com/gofiber/fiber/v2"

	"server/loaders/mysql"
	"server/types/common"
	"server/types/response"

	"share/types/payload"
)

func ActivityGraphHandler(c *fiber.Ctx) error {
	// * Parse session
	s := c.Locals("session").(*common.Session)
	s.SetTag("diary/graph")

	// * Parse query
	query := new(payload.DiaryOverviewQuery)
	if err := c.QueryParser(query); err != nil {
		return response.Error(false, "Unable to parse query", err)
	}

	// * Validate query
	if err := text.Validator.Struct(query); err != nil {
		return response.Error(false, "Unable to validate query", err)
	}

	// * Set session detail
	s.SetDetail("year", query.Year)

	// * Declare variables
	year := *query.Year

	// * Query diary records
	var diaries []*strapiModel.Diary
	if result := mysql.StrapiDB.
		Table("diaries").
		Where("date LIKE ? AND published_at IS NOT NULL", fmt.Sprintf("%d-%%", year)).
		Scan(&diaries); result.Error != nil {
		return response.Error(false, "Unable to query diary records", result.Error)
	}

	// * Construct day mapper
	dayMapper := make(map[string]*strapiModel.Diary)
	for _, diary := range diaries {
		dayMapper[*diary.Date] = diary
	}

	// * Map year to array
	dayGraph := make([]*payload.DiaryDayGraph, 0)
	for iter := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC); iter.Year() == year; iter = iter.AddDate(0, 0, 1) {
		var graph int
		if day, ok := dayMapper[iter.Format("2006-01-02")]; ok {
			var coloredGraph int
			if day.GraphGain != nil {
				coloredGraph += *day.GraphGain * 0xFF / 100
			}
			if day.GraphEmotional != nil {
				coloredGraph += *day.GraphEmotional * 0xFF / 100 << 8
			}
			if day.GraphProductivity != nil {
				coloredGraph += *day.GraphProductivity * 0xFF / 100 << 16
			}
			lightenGraph := (coloredGraph * 0xE0E0E0 / 0xFFFFFF) + 0x101010
			graph = lightenGraph
		} else if iter.Before(time.Now()) {
			graph = 0x121212
		} else {
			graph = 0x181818
		}

		dayGraph = append(dayGraph, &payload.DiaryDayGraph{
			Date:    value.Ptr(iter.Format("2006-01-02")),
			Title:   value.Ptr(iter.Format("Jan 02, 2006")),
			Color:   &graph,
			Summary: nil,
		})
	}

	// * Commit session
	_ = s.Commit(nil)

	return c.JSON(response.Success(map[string]any{
		"days": dayGraph,
	}))
}
