package stateEndpoint

import (
	"backend/generate/sqlc"

	"go.scnd.dev/open/polygon"
	"go.scnd.dev/open/polygon/compat/common"
)

type Handler struct {
	Layer    polygon.Layer
	database common.DatabaseInf[sqlc.Querier]
}

func Handle(
	polygon polygon.Polygon,
	database common.DatabaseInf[sqlc.Querier],
) *Handler {
	return &Handler{
		Layer:    polygon.Layer("state", "endpoint"),
		database: database,
	}
}
