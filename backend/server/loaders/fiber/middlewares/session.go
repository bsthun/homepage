package middlewares

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"share/types/model"
	"share/utils/text"
	value2 "share/utils/value"

	"github.com/gofiber/fiber/v2"

	"server/loaders/mysql"
	"server/types/common"
	"server/types/response"
)

var Session = func() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var session *model.TrackSession
		var newSession bool
		var referenceId *uint64
		if c.Cookies("session_id") == "" {
			newSession = true
		} else {
			if result := mysql.DB.First(&session, c.Cookies("session_id")); result.Error != nil {
				newSession = true
			} else if *session.Token != c.Cookies("session_token") {
				newSession = true
				referenceId = session.Id
			} else if (*session.UpdatedAt).Before(time.Now().Add(-24 * time.Hour)) {
				token := text.Random(text.RandomSet.MixedAlphaNum, 16)
				if result := mysql.DB.Model(session).Update("token", token); result.Error != nil {
					return response.Error(true, "Failed to refresh session token", result.Error)
				}
				session.Token = token
				ApplyCookie(c, session)
			}
		}

		if newSession {
			session = &model.TrackSession{
				Id:         nil,
				Token:      text.Random(text.RandomSet.MixedAlphaNum, 16),
				UserAgent:  value2.Ptr(c.Get("User-Agent")),
				IpAddress:  value2.Ptr(c.Get("X-Real-IP")),
				Attributes: value2.Ptr("{}"),
				CreatedAt:  nil,
				UpdatedAt:  nil,
			}
			if referenceId != nil {
				if bytes, err := json.Marshal(map[string]any{
					"reference_id": referenceId,
				}); err != nil {
					return response.Error(true, "Failed to marshal session attributes", err)
				} else {
					session.Attributes = value2.Ptr(string(bytes))
				}
			}
			if result := mysql.DB.Create(session); result.Error != nil {
				return response.Error(true, "Failed to create session", result.Error)
			}
			ApplyCookie(c, session)
		} else {
			ips := strings.Split(*session.IpAddress, ",")
			if !value2.Contain(ips, c.Get("X-Real-IP")) {
				ips = append(ips, c.Get("X-Real-IP"))
				session.IpAddress = value2.Ptr(strings.Join(ips, ","))
				if result := mysql.DB.Save(session); result.Error != nil {
					return response.Error(true, "Failed to update session", result.Error)
				}
			}
		}

		c.Locals("session", &common.Session{
			Id: *session.Id,
		})

		return c.Next()
	}
}()

func ApplyCookie(c *fiber.Ctx, session *model.TrackSession) {
	c.Cookie(&fiber.Cookie{
		Name:        "session_id",
		Value:       strconv.FormatUint(*session.Id, 10),
		Path:        "/",
		Domain:      "",
		MaxAge:      0,
		Expires:     time.Now().Add(365 * 24 * time.Hour),
		Secure:      false,
		HTTPOnly:    false,
		SameSite:    "",
		SessionOnly: false,
	})
	c.Cookie(&fiber.Cookie{
		Name:        "session_token",
		Value:       *session.Token,
		Path:        "/",
		Domain:      "",
		MaxAge:      0,
		Expires:     time.Now().Add(365 * 24 * time.Hour),
		Secure:      false,
		HTTPOnly:    false,
		SameSite:    "",
		SessionOnly: false,
	})
}
