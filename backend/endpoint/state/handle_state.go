package stateEndpoint

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"go.scnd.dev/open/polygon/compat/predefine"
	"go.scnd.dev/open/polygon/compat/response"
)

func (r *Handler) HandleState(c fiber.Ctx) error {
	// span
	s, _ := r.Layer.With(c.Context())
	defer s.End()

	// * login claims
	var l *predefine.LoginClaims
	if c.Locals("l") != nil {
		l = c.Locals("l").(*jwt.Token).Claims.(*predefine.LoginClaims)
	}

	_ = l

	// * response
	return c.JSON(response.Success(s, map[string]any{
		"state": "ok",
	}))
}
