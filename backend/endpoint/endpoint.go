package endpoint

import (
	"backend/common/config"
	stateEndpoint "backend/endpoint/state"

	"github.com/gofiber/fiber/v3"
	"go.scnd.dev/open/polygon/compat/predefine"
)

func Bind(
	frontend predefine.FrontendFS,
	config *config.Config,
	app *fiber.App,
	stateEndpoint *stateEndpoint.Handler,
) {
	api := app.Group("/api")

	// * state endpoints
	state := api.Group("/state")
	state.Post("/state", stateEndpoint.HandleState)

	// * frontend
	app.Get("*", predefine.EndpointStatic(frontend, config.GetFrontendUrl()))
}
