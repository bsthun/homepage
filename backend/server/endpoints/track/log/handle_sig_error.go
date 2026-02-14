package trackLog

import (
	"github.com/gofiber/fiber/v2"

	"server/types/common"
	"server/types/response"

	"share/types/payload"
)

func SigErrorHandler(c *fiber.Ctx) error {
	// * Parse session
	s := c.Locals("session").(*common.Session)
	s.SetTag("track/sig/error")

	// * Parse body
	body := new(payload.SigBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	s.SetDetail("t", body.Title)

	_ = s.Commit(nil)
	return c.JSON(response.Success(map[string]any{
		"accepted": true,
	}))
}
