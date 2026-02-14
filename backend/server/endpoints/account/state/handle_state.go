package accountState

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"server/types/common"
	"server/types/response"
)

func StateHandler(c *fiber.Ctx) error {
	// * Parse session
	s := c.Locals("session").(*common.Session)
	s.SetTag("account/state")

	if strings.Contains(c.Get("User-Agent"), "wv") || (strings.Contains(c.Get("User-Agent"), "Mobile/") && !strings.Contains(c.Get("User-Agent"), "Safari/")) {
		s.SetDetail("wv", 1)
		_ = s.Commit(nil)
		return response.Error(true, "WEBVIEW_UNLOADED")
	}

	return c.JSON(response.Success(map[string]any{
		"state": "ok",
	}))
}
