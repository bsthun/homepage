package trackLog

import (
	"github.com/gofiber/fiber/v2"

	"server/types/common"
	"server/types/response"

	"share/types/payload"
)

func SectionDetailHandler(c *fiber.Ctx) error {
	// * Parse session
	s := c.Locals("session").(*common.Session)
	s.SetTag("track/log/click")

	// * Parse body
	body := new(payload.LogBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	s.SetDetail("ev", body.Event)
	if body.BeginningState != nil {
		s.SetDetail("begin", body.BeginningState)
	}
	if body.EndingState != nil {
		s.SetDetail("end", body.EndingState)
	}

	_ = s.Commit(nil)
	return c.JSON(response.Success(map[string]any{
		"accepted": true,
	}))
}
