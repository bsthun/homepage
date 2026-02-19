package main

import (
	"backend/common/config"
	"backend/endpoint"
	stateEndpoint "backend/endpoint/state"
	"backend/generate/sqlc"
	"embed"

	"go.scnd.dev/open/polygon/compat/common"
	"go.scnd.dev/open/polygon/compat/predefine"
	"go.uber.org/fx"
)

//go:embed database/migration/*.sql
var migration embed.FS

//go:embed .local/dist/*
var frontend embed.FS

func main() {
	fx.New(
		fx.Supply(
			sqlc.New,
		),
		fx.Provide(
			func() predefine.FrontendFS {
				return frontend
			},
			func() predefine.MigrationFS {
				return migration
			},
			fx.Annotate(
				common.Config[config.Config],
				fx.As(new(common.FiberConfig)),
				fx.As(new(common.DatabaseConfig)),
				fx.As(new(common.PolygonConfig)),
			),
			common.Config[config.Config],
			common.Database[*sqlc.Queries, sqlc.Querier, sqlc.DBTX],
			common.Fiber,
			common.Polygon,
			stateEndpoint.Handle,
		),
		fx.Invoke(
			endpoint.Bind,
		),
	).Run()
}
